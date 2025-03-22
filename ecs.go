package main

import (
	"reflect"
	"slices"
)

type AppState int

const (
	SPLASH_SCREEN AppState = iota
	GAME
	GAME_OVER
)

type ID int
type ComponentMap map[ID]any

type World struct {
	State      AppState
	Ticks      int64
	lastID     ID
	Components map[any]ComponentMap
	Systems    []System
}

type System interface {
	Update(*World)
}

type SystemManager struct {
	Systems []*System
}

func NewWorld() *World {
	return &World{
		State:      SPLASH_SCREEN,
		Ticks:      0,
		lastID:     -1,
		Components: make(map[any]ComponentMap),
	}
}

func (world *World) AllocateID() ID {
	world.lastID += 1
	return world.lastID
}

func (world *World) AddComponent(id ID, component any) {
	componentType := reflect.TypeOf(component)

	if world.Components[componentType] == nil {
		world.Components[componentType] = make(ComponentMap)
	}

	world.Components[componentType][id] = component
}

func (world *World) RemoveComponent(id ID) {
	for _, componentMap := range world.Components {
		if componentMap[id] != nil {
			delete(componentMap, id)
		}
	}
}

func (world *World) Nuke() {
	world.Components = make(map[any]ComponentMap)
	world.Systems = make([]System, 0)
}

func (world *World) AddSystem(system System) {
	world.Systems = append(world.Systems, system)
}

func (world *World) RemoveSystem(system System) {
	for i, s := range world.Systems {
		if s == system {
			world.Systems = slices.Delete(world.Systems, i, 1)
		}
	}
}

func (world *World) UpdateSystems() {
	for _, system := range world.Systems {
		system.Update(world)
	}
}

func ComponentsOfType[T any](world *World) *ComponentMap {
	targetType := reflect.TypeOf((*T)(nil))

	for componentType, componentMap := range world.Components {
		if componentType == targetType {
			return &componentMap
		}
	}
	return &ComponentMap{}
}

func Map1[T any](world *World, f func(id ID, componentA *T)) {
	for id, component := range *ComponentsOfType[T](world) {
		f(id, component.(*T))
	}
}

func Map2[T1 any, T2 any](world *World, f func(id ID, componentA *T1, componentB *T2)) {
	componentsA := ComponentsOfType[T1](world)
	componentsB := ComponentsOfType[T2](world)

	// Eww!
	for idA, componentA := range *componentsA {
		if componentB, ok := (*componentsB)[idA]; ok {
			f(idA, componentA.(*T1), componentB.(*T2))
		}
	}
}

func Map3[T1 any, T2 any, T3 any](world *World, f func(id ID, componentA *T1, componentB *T2, componentC *T3)) {
	componentsA := ComponentsOfType[T1](world)
	componentsB := ComponentsOfType[T2](world)
	componentsC := ComponentsOfType[T3](world)

	// Eww! Eww!
	for idA, componentA := range *componentsA {
		if componentB, ok := (*componentsB)[idA]; ok {
			if componentC, ok := (*componentsC)[idA]; ok {
				f(idA, componentA.(*T1), componentB.(*T2), componentC.(*T3))
			}
		}
	}
}

func Map4[T1 any, T2 any, T3 any, T4 any](world *World, f func(id ID, componentA *T1, componentB *T2, componentC *T3, componentD *T4)) {
	componentsA := ComponentsOfType[T1](world)
	componentsB := ComponentsOfType[T2](world)
	componentsC := ComponentsOfType[T3](world)
	componentsD := ComponentsOfType[T4](world)

	// Eww! Eww! Eww!
	for idA, componentA := range *componentsA {
		if componentB, ok := (*componentsB)[idA]; ok {
			if componentC, ok := (*componentsC)[idA]; ok {
				if componentD, ok := (*componentsD)[idA]; ok {
					f(idA, componentA.(*T1), componentB.(*T2), componentC.(*T3), componentD.(*T4))
				}
			}
		}
	}
}
