package main

type Tree struct {
	Positive *Tree
	Value    string
	Negative *Tree
}

func (t *Tree) IsLeaf() bool {
	return t.Positive == nil && t.Negative == nil
}

func (t *Tree) Next(positive bool) (*Tree, string) {
	if t.IsLeaf() {
		return t, ""
	}
	if positive {
		return t.Positive, t.Positive.Value
	}
	return t.Negative, t.Negative.Value
}

func GetAnimals(root *Tree) []string {
	if root == nil {
		return []string{}
	}

	queue := make([]*Tree, 0)
	animals := make([]string, 0)

	queue = append(queue, root)

	for !(len(queue) == 0) {
		temp := queue[0]
		queue = queue[1:]

		if temp.IsLeaf() {
			animals = append(animals, temp.Value)
		}

		if temp.Positive != nil {
			queue = append(queue, temp.Positive)
		}
		if temp.Negative != nil {
			queue = append(queue, temp.Negative)
		}
	}
	return animals
}

func insert(positive *Tree, negative *Tree, v string) *Tree {
	return &Tree{positive, v, negative}
}

func NewAnimalTree() *Tree {
	cat := insert(nil, nil, "cat")
	dog := insert(nil, nil, "dog")
	horse := insert(nil, nil, "horse")
	whale := insert(nil, nil, "whale")
	elephant := insert(nil, nil, "elephant")
	zebra := insert(nil, nil, "zebra")
	wolf := insert(nil, nil, "wolf")
	deer := insert(nil, nil, "deer")
	eagle := insert(nil, nil, "eagle")
	parrot := insert(nil, nil, "parrot")
	chicken := insert(nil, nil, "chicken")
	penguin := insert(nil, nil, "penguin")
	shark := insert(nil, nil, "shark")
	spider := insert(nil, nil, "spider")
	cobra := insert(nil, nil, "cobra")
	hypo := insert(nil, nil, "hippopotamus")
	pigeon := insert(nil, nil, "pigeon")
	catfish := insert(nil, nil, "catfish")
	tuna := insert(nil, nil, "tuna")
	tentacles := insert(spider, chicken, "Does it have tentacles?")
	salt := insert(catfish, tuna, "Does it live in fresh water?")
	dangerous := insert(shark, salt, "Can it eat humans?")
	venom := insert(cobra, dangerous, "Does it use venom to hunt?")
	legs := insert(tentacles, venom, "Does it have legs?")
	cold := insert(penguin, legs, "Can it leave under extreme cold conditions?")
	america := insert(eagle, pigeon, "Is it one of the symbol of the USA?")
	human := insert(parrot, america, "Can it speak human language?")
	fly := insert(human, cold, "Does it fly?")
	trunk := insert(elephant, hypo, "Does it have a trunk?")
	horns := insert(deer, wolf, "Does the male have horns?")
	heavy := insert(trunk, horns, "Is it gennerally as heavy as a truck?")
	stripped := insert(zebra, horse, "Is it stripped?")
	galop := insert(stripped, heavy, "Does it gallop?")
	bark := insert(dog, cat, "Does it bark and growl?")
	humans := insert(bark, galop, "Is it generally living with humans?")
	water := insert(whale, humans, "Does it live in the water?")
	mammal := insert(water, fly, "Is it a mammal?")
	root := insert(mammal, nil, "")
	return root
}
