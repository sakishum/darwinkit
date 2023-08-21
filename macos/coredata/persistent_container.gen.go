// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentContainer] class.
var PersistentContainerClass = _PersistentContainerClass{objc.GetClass("NSPersistentContainer")}

type _PersistentContainerClass struct {
	objc.Class
}

// An interface definition for the [PersistentContainer] class.
type IPersistentContainer interface {
	objc.IObject
	NewBackgroundContext() ManagedObjectContext
	LoadPersistentStoresWithCompletionHandler(block func(arg0 PersistentStoreDescription, arg1 foundation.Error))
	PerformBackgroundTask(block func(arg0 ManagedObjectContext))
	PersistentStoreDescriptions() []PersistentStoreDescription
	SetPersistentStoreDescriptions(value []IPersistentStoreDescription)
	ManagedObjectModel() ManagedObjectModel
	PersistentStoreCoordinator() PersistentStoreCoordinator
	Name() string
	ViewContext() ManagedObjectContext
}

// A container that encapsulates the Core Data stack in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer?language=objc
type PersistentContainer struct {
	objc.Object
}

func PersistentContainerFrom(ptr unsafe.Pointer) PersistentContainer {
	return PersistentContainer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PersistentContainer) InitWithName(name string) PersistentContainer {
	rv := objc.Call[PersistentContainer](p_, objc.Sel("initWithName:"), name)
	return rv
}

// Creates a container with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640557-initwithname?language=objc
func NewPersistentContainerWithName(name string) PersistentContainer {
	instance := PersistentContainerClass.Alloc().InitWithName(name)
	instance.Autorelease()
	return instance
}

func (pc _PersistentContainerClass) PersistentContainerWithName(name string) PersistentContainer {
	rv := objc.Call[PersistentContainer](pc, objc.Sel("persistentContainerWithName:"), name)
	return rv
}

// Initializes a new persistent container using the provided name for the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1646295-persistentcontainerwithname?language=objc
func PersistentContainer_PersistentContainerWithName(name string) PersistentContainer {
	return PersistentContainerClass.PersistentContainerWithName(name)
}

func (pc _PersistentContainerClass) Alloc() PersistentContainer {
	rv := objc.Call[PersistentContainer](pc, objc.Sel("alloc"))
	return rv
}

func PersistentContainer_Alloc() PersistentContainer {
	return PersistentContainerClass.Alloc()
}

func (pc _PersistentContainerClass) New() PersistentContainer {
	rv := objc.Call[PersistentContainer](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentContainer() PersistentContainer {
	return PersistentContainerClass.New()
}

func (p_ PersistentContainer) Init() PersistentContainer {
	rv := objc.Call[PersistentContainer](p_, objc.Sel("init"))
	return rv
}

// Returns a new managed object context that executes on a private queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640581-newbackgroundcontext?language=objc
func (p_ PersistentContainer) NewBackgroundContext() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](p_, objc.Sel("newBackgroundContext"))
	rv.Autorelease()
	return rv
}

// Loads the persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640568-loadpersistentstoreswithcompleti?language=objc
func (p_ PersistentContainer) LoadPersistentStoresWithCompletionHandler(block func(arg0 PersistentStoreDescription, arg1 foundation.Error)) {
	objc.Call[objc.Void](p_, objc.Sel("loadPersistentStoresWithCompletionHandler:"), block)
}

// Returns the location of the directory that contains the persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640597-defaultdirectoryurl?language=objc
func (pc _PersistentContainerClass) DefaultDirectoryURL() foundation.URL {
	rv := objc.Call[foundation.URL](pc, objc.Sel("defaultDirectoryURL"))
	return rv
}

// Returns the location of the directory that contains the persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640597-defaultdirectoryurl?language=objc
func PersistentContainer_DefaultDirectoryURL() foundation.URL {
	return PersistentContainerClass.DefaultDirectoryURL()
}

// Executes a block on a private queue using an ephemeral managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640564-performbackgroundtask?language=objc
func (p_ PersistentContainer) PerformBackgroundTask(block func(arg0 ManagedObjectContext)) {
	objc.Call[objc.Void](p_, objc.Sel("performBackgroundTask:"), block)
}

// The descriptions of the container’s persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640577-persistentstoredescriptions?language=objc
func (p_ PersistentContainer) PersistentStoreDescriptions() []PersistentStoreDescription {
	rv := objc.Call[[]PersistentStoreDescription](p_, objc.Sel("persistentStoreDescriptions"))
	return rv
}

// The descriptions of the container’s persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640577-persistentstoredescriptions?language=objc
func (p_ PersistentContainer) SetPersistentStoreDescriptions(value []IPersistentStoreDescription) {
	objc.Call[objc.Void](p_, objc.Sel("setPersistentStoreDescriptions:"), value)
}

// The container’s managed object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640561-managedobjectmodel?language=objc
func (p_ PersistentContainer) ManagedObjectModel() ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](p_, objc.Sel("managedObjectModel"))
	return rv
}

// The container’s persistent store coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640567-persistentstorecoordinator?language=objc
func (p_ PersistentContainer) PersistentStoreCoordinator() PersistentStoreCoordinator {
	rv := objc.Call[PersistentStoreCoordinator](p_, objc.Sel("persistentStoreCoordinator"))
	return rv
}

// The container’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640579-name?language=objc
func (p_ PersistentContainer) Name() string {
	rv := objc.Call[string](p_, objc.Sel("name"))
	return rv
}

// The main queue’s managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640622-viewcontext?language=objc
func (p_ PersistentContainer) ViewContext() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](p_, objc.Sel("viewContext"))
	return rv
}