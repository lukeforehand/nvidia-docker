// Copyright (c) 2015, NVIDIA CORPORATION. All rights reserved.

package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path"

	"graceful"
)

type plugin interface {
	implement() string
	register(*PluginAPI)
}

type PluginAPI struct {
	*graceful.HTTPServer

	plugins []plugin
}

func accept(handler http.Handler) http.Handler {
	const header = "application/vnd.docker.plugins.v1.1+json"

	f := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") != header {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}

func NewPluginAPI(addr string) *PluginAPI {
	os.MkdirAll(path.Dir(addr), 0700)

	a := &PluginAPI{
		HTTPServer: graceful.NewHTTPServer("unix", addr, accept),
	}
	a.Handle("POST", "/Plugin.Activate", a.activate)

	a.register(
		new(pluginVolume),
	)
	return a
}

func (a *PluginAPI) register(plugins ...plugin) {
	for _, p := range plugins {
		p.register(a)
		a.plugins = append(a.plugins, p)
	}
}

func (a *PluginAPI) activate(resp http.ResponseWriter, req *http.Request) {
	r := struct{ Implements []string }{}

	r.Implements = make([]string, len(a.plugins))
	for i, p := range a.plugins {
		r.Implements[i] = p.implement()
	}
	assert(json.NewEncoder(resp).Encode(r))
}
