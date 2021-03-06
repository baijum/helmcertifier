/*
 * Copyright (C) 06/01/2021, 09:40, igors
 * This file is part of helmcertifier.
 *
 * helmcertifier is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * helmcertifier is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with helmcertifier.  If not, see <http://www.gnu.org/licenses/>.
 */

package checks

type Result struct {
	Ok bool
}

type CheckFunc func(uri string) (Result, error)

type Registry interface {
	Get(name string) (CheckFunc, bool)
	Add(name string, checkFunc CheckFunc) Registry
}

type defaultRegistry map[string]CheckFunc

func NewRegistry() Registry {
	return &defaultRegistry{}
}

func (r *defaultRegistry) Get(name string) (CheckFunc, bool) {
	v, ok := (*r)[name]
	return v, ok
}

func (r *defaultRegistry) Add(name string, checkFunc CheckFunc) Registry {
	(*r)[name] = checkFunc
	return nil
}
