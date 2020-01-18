package xconfig

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type (
	Registry struct {
		SecretKey        string
		configNamespaces [][]string
		configs          []interface{}
		loaded           bool
		doOnReady        []func()
	}
)

func NewRegistry(secretKey string) *Registry {
	return &Registry{
		SecretKey:        secretKey,
		configs:          []interface{}{},
		configNamespaces: [][]string{},
		doOnReady:        []func(){},
		loaded:           false,
	}
}

func (c *Registry) Register(o interface{}, namespace ...string) {
	c.configs = append(c.configs, o)
	c.configNamespaces = append(c.configNamespaces, namespace)
}

func (c *Registry) Load(path string) (err error) {
	defer func() { c.loaded = true }()

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	node := new(yaml.Node)
	err = yaml.Unmarshal(content, node)
	if err == nil {
		err = c.parseConfigs(node)
	}
	if err == nil {
		for _, v := range c.doOnReady {
			v()
		}
	}
	return
}

func (c *Registry) DoOnReady(f func()) {
	if c.loaded {
		f()
	} else {
		c.doOnReady = append(c.doOnReady, f)
	}
}

func (c *Registry) parseConfigs(rootNode *yaml.Node) error {
	for i, v := range c.configs {
		node, err := getSubNode(rootNode, c.configNamespaces[i])
		if err != nil {
			return err
		}
		err = node.Decode(v)
		if err != nil {
			return err
		} else if c.SecretKey != "" {
			err = Decode(v, c.SecretKey)
		}
	}
	return nil
}
func getSubNode(node *yaml.Node, namespace []string) (*yaml.Node, error) {
	if len(namespace) == 0 {
		return node, nil
	}
	if node.Kind == yaml.DocumentNode {
		return getSubNode(node.Content[0], namespace)
	} else if node.Kind == yaml.MappingNode {
		for i, v := range node.Content {
			if v.Kind == yaml.ScalarNode && v.Value == namespace[0] {
				return getSubNode(node.Content[i+1], namespace[1:])
			}
		}
		return nil, fmt.Errorf("Can't find the key '%s' in map at line %d", namespace[0], node.Line)
	} else {
		return nil, fmt.Errorf("Can't find namespace %s because the node is %s at line %d", namespace, node.Kind, node.Line)
	}
}
