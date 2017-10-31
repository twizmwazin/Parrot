package world

//represents the contain of the world.
//allows you to obtain different chunks of world
type World interface {
	//getChunkAt returns a chuck at the given coordinates
	//if no chunk will create one at given coordinates
	getChunkAt(x int, z int) Chunk

	//TODO: more general world methods needed
}

//Chunk repersents chunks within the game
//when changed will change the actual mbt files
//of the game with specified changes
type Chunk interface {
	//getBlockAt returns a block at the given coordinates
	//if no block at location will error saying no block at location
	getBlockAt(x int, y int, z int) Block

	//updateBlock adds a given block to map or updates block currently there
	//returns a bool on weather it was added correctly
	updateBlock(block Block) bool

	//getItems returns a list of items in the chunk
	getItems() []Item

	//updateItem will update item at location or place it there
	//returns a bull telling weather it was added correctly
	updateItem(item Item) bool

	//placeItem places the item in the current Chunk
	//returns if it was successfully added
	placeItem(item Item) bool

	//getLocation returns the location of the Chunk
	getLocation() (int, int)
}

//repersents the blocks of map
type Block interface {
	//getLocation returns the location of the Block
	getLocation() (int, int, int)

	//TODO: change type of get type
	//returns the type of the block
	getType() string

	//changeType takes in a type to update the block with
	//returns if it was successfully updated
	changeType(t string) bool
}

//might not be perfect name
//represents all other things that are not blocks
//meant to be generic to all other types of objects
type Item interface {
	//getLocation returns the location of the Item
	getLocation() (float32, float32, float32)

	//TODO: change type of get type
	//returns the type of the Item
	getType() string
}
