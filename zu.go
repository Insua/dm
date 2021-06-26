/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm

import (
	"strconv"
	"strings"
	"sync"
)

type Properties struct {
	innerProps map[string]string
	sync.RWMutex
}

func NewProperties() *Properties {
	p := Properties{
		innerProps: make(map[string]string, 50),
	}
	return &p
}

func (g *Properties) SetProperties(p *Properties) {
	if p == nil {
		return
	}
	for k, v := range p.innerProps {
		g.Set(strings.ToLower(k), v)
	}
}

func (g *Properties) Len() int {
	return len(g.innerProps)
}

func (g *Properties) IsNil() bool {
	return g == nil || g.innerProps == nil
}

func (g *Properties) GetString(key, def string) string {
	g.RLock()
	v, ok := g.innerProps[strings.ToLower(key)]
	g.RUnlock()

	if !ok || v == "" {
		return def
	}
	return v
}

func (g *Properties) GetInt(key string, def int, min int, max int) int {
	g.RLock()
	value, ok := g.innerProps[strings.ToLower(key)]
	g.RUnlock()
	if !ok || value == "" {
		return def
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	}

	if i > max || i < min {
		return def
	}
	return i
}

func (g *Properties) GetBool(key string, def bool) bool {
	g.RLock()
	value, ok := g.innerProps[strings.ToLower(key)]
	g.RUnlock()
	if !ok || value == "" {
		return def
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}
	return b
}

func (g *Properties) GetTrimString(key string, def string) string {
	g.RLock()
	value, ok := g.innerProps[strings.ToLower(key)]
	g.RUnlock()
	if !ok || value == "" {
		return def
	} else {
		return strings.TrimSpace(value)
	}
}

func (g *Properties) GetStringArray(key string, def []string) []string {
	g.RLock()
	value, ok := g.innerProps[strings.ToLower(key)]
	g.RUnlock()
	if ok || value != "" {
		array := strings.Split(value, ",")
		if len(array) > 0 {
			return array
		}
	}
	return def
}

//func (g *Properties) GetBool(key string) bool {
//	i, _ := strconv.ParseBool(g.innerProps[key])
//	return i
//}

func (g *Properties) Set(key, value string) {
	g.Lock()
	g.innerProps[strings.ToLower(key)] = value
	g.Unlock()
}

// 如果p有g没有的键值对,添加进g中
func (g *Properties) SetDiffProperties(p *Properties) {
	if p == nil {
		return
	}
	g.Lock()
	for k, v := range p.innerProps {
		if _, ok := g.innerProps[strings.ToLower(k)]; !ok {
			g.innerProps[strings.ToLower(k)] = v
		}
	}
	g.Unlock()
}
