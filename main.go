package main

import (
	"fmt"
	"strconv"
	"strings"
)

// обозначаем структуру блюдо
type dish struct {
	nameDish         string //назван блюда
	peopleEat        int    //количество людей, желающих отведать блюдо
	numberIngredient int    // сколько ингридентов в блюде
}

// создание нового блюда
func newDish(l string) *dish {
	s := strings.Split(l, " ")
	d := dish{}
	d.nameDish = s[0]
	d.peopleEat, _ = strconv.Atoi(s[1])
	d.numberIngredient, _ = strconv.Atoi(s[2])
	return &d
}

// хранилище блюд
type dishStorage struct {
	dishName map[string]*dish
}

// создание хранилища блюд
func newDishStorage() *dishStorage {
	return &dishStorage{dishName: make(map[string]*dish)}

}

// добавление блюда в хранилище
func (ds *dishStorage) put(d *dish) {
	ds.dishName[d.nameDish] = d

}
func (ds *dishStorage) getAll() []*dish {
	var dishs []*dish
	for _, v := range ds.dishName {
		dishs = append(dishs, v)
	}
	return dishs
}

// обозначаем структуру ингредиент
type ingredient struct {
	nameIng   string // название ингредиента
	amountIng int    //количество игредиента
	unit      string //название идиницы измерения
}

// создание нового ингредиента
func newIngredient(l string) *ingredient {
	s := strings.Split(l, " ")
	i := ingredient{}
	i.nameIng = s[0]
	i.amountIng, _ = strconv.Atoi(s[1])
	i.unit = s[2]
	return &i
}
func newIngredientEat(i *ingredient, k int) *ingredient {
	ing := ingredient{}
	ing.nameIng = i.nameIng
	ing.amountIng = i.amountIng * k
	ing.unit = i.unit
	return &ing

}

type ingredientStorage struct {
	ingredientName map[string]*ingredient
}

func newIngredientStorage() *ingredientStorage {
	return &ingredientStorage{make(map[string]*ingredient)}
}

func (is *ingredientStorage) put(i *ingredient) {
	is.ingredientName[i.nameIng] = i
}
func (is *ingredientStorage) getAll() []*ingredient {
	var ingredients []*ingredient
	for _, v := range is.ingredientName {
		ingredients = append(ingredients, v)
	}
	return ingredients
}

type calories struct {
	nameIng string  // — название ингредиента, состоящее из строчных букв английского алфавита, цифр и знака подчёркивания. Длина строки не превосходит 20 символов;
	ai      int     // — количество ингредиента, для которого указаны характеристики ингредиента, заданное целым числом (1⩽ai⩽1000);
	ui      string  //— единица измерения (l, ml, g, kg, cnt или tens);
	pri     float64 // содержание белков, шесть знаков после запятой
	fi      float64 // содержание жиров
	chi     float64 // содержание углеводов
	fvi     float64 // энергетическая ценность ингредиента, соответственно,
}

func newCalories(l string) *calories {
	s := strings.Split(l, " ")
	c := calories{}
	c.nameIng = s[0]
	c.ai, _ = strconv.Atoi(s[1])
	c.ui = s[2]
	c.pri, _ = strconv.ParseFloat(s[3], 64)
	c.fi, _ = strconv.ParseFloat(s[4], 64)
	c.chi, _ = strconv.ParseFloat(s[5], 64)
	c.fvi, _ = strconv.ParseFloat(s[6], 64)
	return &c
}

type caloriesStorage struct {
	ingredientName map[string]*calories
}

func newIngredientCaloriesStorage() *caloriesStorage {
	return &caloriesStorage{make(map[string]*calories)}
}
func (ics *caloriesStorage) put(i *calories) {
	ics.ingredientName[i.nameIng] = i
}
func (ics *caloriesStorage) getAll() []*calories {
	var ingredients []*calories
	for _, v := range ics.ingredientName {
		ingredients = append(ingredients, v)
	}
	return ingredients
}

type caloriesDish struct {
	pri float64 // содержание белков, шесть знаков после запятой
	fi  float64 // содержание жиров
	chi float64 // содержание углеводов
	fvi float64 // энергетическая ценность ингредиента, соответственно,
}
type basket struct {
	nameIngredient  string //Название ингредиента
	priceIngredient int    //стоимость ингредиента
	ai              int    //количество ингредиента
	ui              string // единица измерения
}

func newBasket(l string) *basket {
	s := strings.Split(l, " ")
	b := basket{}
	b.nameIngredient = s[0]
	b.priceIngredient, _ = strconv.Atoi(s[1])
	b.ai, _ = strconv.Atoi(s[2])
	b.ui = s[3]
	return &b
}
func newIngredientBas(i *ingredient, b *basket) *basket {
	unit := 1
	if b.ui == "l" || b.ui == "kg" {
		unit = 1000
	}
	if b.ui == "tens" {
		unit = 10
	}
	x := i.amountIng / (b.ai * unit)
	if i.amountIng/b.ai != 0 || x == 0 {
		x++
	}

	bas := basket{}
	bas.nameIngredient = i.nameIng // имя ингредиента
	bas.ui = b.ui
	bas.ai = x                                       //единица измерения
	bas.priceIngredient = b.priceIngredient * bas.ai //итоговая цена

	return &bas

}

type basketStorage struct {
	basketName map[string]*basket
}

func newBasketStorage() *basketStorage {
	return &basketStorage{basketName: make(map[string]*basket)}
}
func (bs *basketStorage) put(b *basket) {
	bs.basketName[b.nameIngredient] = b
}
func (bs *basketStorage) getAll() []*basket {
	var baskets []*basket
	for _, v := range bs.basketName {
		baskets = append(baskets, v)
	}
	return baskets
}

type ingredientMap struct {
	inredientStorageName map[int]*ingredientStorage
}

func newIngredientMap() *ingredientMap {
	return &ingredientMap{inredientStorageName: make(map[int]*ingredientStorage)}
}
func (im *ingredientMap) put(is *ingredientStorage, i int) {
	im.inredientStorageName[i] = is
}
func (im *ingredientMap) getAll() []*ingredientStorage {
	var is []*ingredientStorage
	for _, v := range im.inredientStorageName {
		is = append(is, v)
	}
	return is
}

func main() {
	ds := newDishStorage()
	ics := newIngredientCaloriesStorage()
	im := newIngredientMap()
	ingredientMapEat := newIngredientMap()
	bs := newBasketStorage()
	s := "2\n" + //сколько блюд готовить
		"sandwich 7 3\n" + //название блюдо, на сколько человек, количество ингредиентов
		"butter 10 g\n" + //ингредиент, сколько гр, мл, шт ;
		"toasted_bread 2 cnt\n" + //ингредиент, сколько гр, мл, шт ;
		"sausage 30 g\n" + //ингредиент, сколько гр, мл, шт ;
		"omelet 9 4\n" + //название блюдо, на сколько человек, количество ингредиентов
		"egg 4 cnt\n" + //ингредиент, сколько гр, мл, шт ;
		"milk 120 ml\n" + //ингредиент, сколько гр, мл, шт ;
		"salt 1 g\n" + //ингредиент, сколько гр, мл, шт ;
		"sausage 50 g\n" + //+ //ингредиент, сколько гр, мл, шт ;
		"7\n" + // количество ингредиентов в магазине;
		"egg 61 1 tens\n" + // ингредиент, цена  за сколько гр, мл, шт ;
		"milk 58 1 l\n" + // ингредиент, цена  за сколько гр, мл, шт ;
		"sausage 100 480 g\n" + // ингредиент, цена  за сколько гр, мл, шт ;
		"butter 120 180 g\n" + // ингредиент, цена  за сколько гр, мл, шт ;
		"cream 100 350 g\n" + // ингредиент, цена  за сколько гр, мл, шт ;
		"salt 14 1000 g\n" + // ингредиент, цена  за сколько гр, мл, шт ;
		"toasted_bread 40 20 cnt\n" + // ингредиент, цена  за сколько гр, мл, шт ;
		"8\n" + // количество ингредиентов с калорийностью
		"egg 1 cnt 13 12 1 16.4\n" + // ингредиент, единица измерения, бжу, ккал
		"milk 1 l 3 4.5 4.7 60\n" + //ингредиент, единица измерения, бжу, ккал
		"chocolate 90 g 6.8 36.3 47.1 546\n" + // ингредиент, единица измерения, бжу, ккал
		"salt 1 kg 0 0 0 0\n" + // ингредиент, единица измерения, бжу, ккал
		"strawberry 100 g 0.4 0.1 7 35\n" + // ингредиент, единица измерения, бжу, ккал
		"sausage 100 g 10 18 1.5 210\n" + // ингредиент, единица измерения, бжу, ккал
		"toasted_bread 5 cnt 7.3 1.6 52.3 248\n" + // ингредиент, единица измерения, бжу, ккал
		"butter 100 g 0.8 72.5 1.3 661 " // ингредиент, единица измерения, бжу, ккал
	s2 := strings.Split(s, "\n")
	i := 0
	amountBush := stringToInt(s2[i])
	i++
	for k := 0; k < amountBush; k++ {
		d := newDish(s2[i])
		i++
		ds.put(d)
		is := newIngredientStorage()
		for j := 0; j < d.numberIngredient; j++ {
			ing := newIngredient(s2[i])
			is.put(ing)
			i++
		}
		im.put(is, k)
	}
	basketInt := stringToInt(s2[i])
	i++
	for x := 0; x < basketInt; x++ {
		bas := newBasket(s2[i])
		bs.put(bas)
		i++
	}
	caloriesInt := stringToInt(s2[i])
	i++
	for y := 0; y < caloriesInt; y++ {
		cal := newCalories(s2[i])
		i++
		ics.put(cal)
	}
	sandvich, omlet := caloriesDish{}, caloriesDish{}
	for i, i3 := range im.getAll() {
		if i == 0 {
			sandvich = totalCalories(*i3, *ics)
		}
		if i == 1 {
			omlet = totalCalories(*i3, *ics)
		}
	}

	for p, value := range ds.getAll() {
		is := newIngredientStorage()
		for s3, storage := range im.getAll() {
			if p == s3 {
				is = totalIngredients(value, storage)
				ingredientMapEat.put(is, p)
			}
		}
	}
	for _, storage := range ingredientMapEat.getAll() {
		for _, i3 := range storage.getAll() {
			fmt.Println(i3)
		}
	}
	fmt.Println()

	tsi := totalShopIngredient(ingredientMapEat)

	for _, i2 := range tsi.getAll() {
		fmt.Println(i2)
	}

	fmt.Println()
	sum := 0
	tbi := totalShopPac(tsi, bs)
	for _, i2 := range tbi.getAll() {
		sum += i2.priceIngredient
	}
	fmt.Println(sum)

	for _, i3 := range tbi.getAll() {
		fmt.Printf("%s %v \n", i3.nameIngredient, i3.ai)
	}
	fmt.Printf("sandvich %.6v %.6v %.6v %.6v\n", sandvich.pri, sandvich.fvi, sandvich.fi, sandvich.chi)
	fmt.Printf("sandvich %.6v %.6v %.6v %.6v\n", omlet.pri, omlet.fvi, omlet.fi, omlet.chi)
}
func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func countingCalories(i ingredient, c calories) caloriesDish {
	ml := c.ai
	var cd caloriesDish
	if c.ui == "l" || c.ui == "kg" {
		ml = 1000
	}
	if c.ui == "tens" {
		ml = 10
	}
	cd.pri = c.pri / float64(ml) * float64(i.amountIng)
	cd.chi = c.chi / float64(ml) * float64(i.amountIng)
	cd.fi = c.fi / float64(ml) * float64(i.amountIng)
	cd.fvi = c.fvi / float64(ml) * float64(i.amountIng)
	return cd
}
func sumCalories(s caloriesDish, s2 caloriesDish) caloriesDish {
	s.chi = s.chi + s2.chi
	s.fi = s.fi + s2.fi
	s.pri = s.pri + s2.pri
	s.fvi = s.fvi + s2.fvi
	return s
}
func totalCalories(is ingredientStorage, cs caloriesStorage) caloriesDish {
	var cd caloriesDish
	for _, val := range is.getAll() {
		for _, val2 := range cs.getAll() {
			if val.nameIng == val2.nameIng {
				cc := countingCalories(*val, *val2)
				cd = sumCalories(cd, cc)
				break
			}
		}
	}
	return cd
}
func totalIngredients(d *dish, is *ingredientStorage) *ingredientStorage {

	ise := newIngredientStorage()
	for _, val := range is.getAll() {
		ing := newIngredientEat(val, d.peopleEat)
		ise.put(ing)
	}
	return ise
}
func totalShopIngredient(im *ingredientMap) *ingredientStorage {
	is := newIngredientStorage()

	for k, storage := range im.getAll() {
		for _, i2 := range storage.getAll() {
			if k < 1 {
				is.put(i2)
			} else {
				for _, i3 := range is.getAll() {
					if i3.nameIng == i2.nameIng {
						i3.amountIng = i3.amountIng + i2.amountIng
						is.put(i3)
						break
					} else {
						is.put(i2)
					}
				}
			}
		}
	}
	return is
}
func totalShopPac(is *ingredientStorage, bs *basketStorage) *basketStorage {
	ip := *newBasketStorage()
	for _, val := range is.getAll() {
		for _, val2 := range bs.getAll() {
			if val.nameIng == val2.nameIngredient {
				nib := newIngredientBas(val, val2)
				ip.put(nib)
			}
		}
	}
	return &ip
}
