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

// хранилище ингредиентов в блюде
type ingredientStorage struct {
	ingredientName map[string]*ingredient
}

// создание хранилища ингредиентов
func newIngredientStorage() *ingredientStorage {
	return &ingredientStorage{make(map[string]*ingredient)}
}

// добавление ингредиента в хранилище
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

type basket struct {
	ti string //Название ингредиента
	pi int    //стоимость ингредиента
	ai int    //количество ингредиента
	ui string // единица измерения
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
type caloriesDish struct {
	pri float64 // содержание белков, шесть знаков после запятой
	fi  float64 // содержание жиров
	chi float64 // содержание углеводов
	fvi float64 // энергетическая ценность ингредиента, соответственно,
}

func main() {
	ds := newDishStorage()
	ingredientMap := make(map[int]*ingredientStorage)
	//basketMap := make(map[int]basket)

	//caloriesMap := make(map[int]calories)*/

	s := "2\n" + //сколько блюд готовить
		"sandwich 7 3\n" + //название блюдо, на сколько человек, количество ингредиентов
		"butter 10 g\n" + //ингредиент, сколько гр, мл, шт ;
		"toasted_bread 2 cnt\n" + //ингредиент, сколько гр, мл, шт ;
		"sausage 30 g\n" + //ингредиент, сколько гр, мл, шт ;
		"omelet 9 4\n" + //название блюдо, на сколько человек, количество ингредиентов
		"egg 4 cnt\n" + //ингредиент, сколько гр, мл, шт ;
		"milk 120 ml\n" + //ингредиент, сколько гр, мл, шт ;
		"salt 1 g\n" + //ингредиент, сколько гр, мл, шт ;
		"sausage 50 g\n" //+ //ингредиент, сколько гр, мл, шт ;
	s2 := strings.Split(s, "\n")
	i := 0
	amountBush := stringToInt(s2[i])
	i++
	for k := 1; k <= amountBush; k++ {
		fmt.Printf("берем строку %s и переделываем в мапу типа dish \n", s2[i])
		d := newDish(s2[i])
		fmt.Println(s2[i])
		fmt.Println(d)
		i++
		fmt.Printf("%v блюдо это %v \n", k, d)
		ds.put(d)
		is := newIngredientStorage()
		for j := 0; j < d.numberIngredient; j++ {
			fmt.Printf("берем строку %s и переделываем в мапу типа ingredient \n", s2[i])
			ing := newIngredient(s2[i])
			is.put(ing)
			i++
		}
		ingredientMap[k] = is
	}
	for _, v := range ds.getAll() {
		fmt.Println(v)
	}
	for _, storage := range ingredientMap {
		for _, v := range storage.getAll() {
			fmt.Println(v)
		}
	}

}

/*"7\n" + // количество ингредиентов в магазине;
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
"  butter 100 g 0.8 72.5 1.3 661 " // ингредиент, единица измерения, бжу, ккал

*/
/*var students []*DishIngredient
for _, v := range DishIngredientMap {
	students = append(students, v)
}
fmt.Println(students)*/
/*basketInt := stringToInt(s2[i])
fmt.Printf("количество продуктов в магазине %v \n", basketInt)
i++
for x := 0; x < basketInt; x++ {
	bas := stringToBasket(s2[i])
	basketMap[x] = bas
	i++
}
fmt.Println("получили список продуктов в магазине")
fmt.Println(basketMap)
caloriesInt := stringToInt(s2[i])
fmt.Printf("количество продуктов c калорийностью %v \n", caloriesInt)
i++
for y := 0; y < caloriesInt; y++ {
	cal := stringToCalories(s2[i])
	caloriesMap[y] = cal
	i++
}
fmt.Println("создали хранилище с калорийностью продуктов")
fmt.Println(caloriesMap)
var caloriesSandwich caloriesDish
fmt.Println(ingredientMap)
for _, val := range ingredientMap[0] {
	fmt.Printf("берем первое значение %v \n", val)
	for _, val2 := range caloriesMap {
		fmt.Printf("берем второе значение %v \n", val2)
		if val.nameIng == val2.nameIng {
			cc := countingCalories(val, val2)
			caloriesSandwich = sumCalories(caloriesSandwich, cc)
			break
		}

	}
}
fmt.Println("подсчитали калорийность сэндвича")
fmt.Println(caloriesSandwich)*/

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func stringToFloat(s string) float64 {
	result, _ := strconv.ParseFloat(s, 64)
	return result
}

func stringToBasket(s string) basket {
	s2 := strings.Split(s, " ")
	d := basket{
		ti: s2[0],
		pi: stringToInt(s2[1]),
		ai: stringToInt(s2[2]),
		ui: s2[3],
	}
	return d
}
func stringToCalories(s string) calories {
	s2 := strings.Split(s, " ")
	c := calories{
		nameIng: s2[0],                // — название ингредиента, состоящее из строчных букв английского алфавита, цифр и знака подчёркивания. Длина строки не превосходит 20 символов;
		ai:      stringToInt(s2[1]),   // — количество ингредиента, для которого указаны характеристики ингредиента, заданное целым числом (1⩽ai⩽1000);
		ui:      s2[2],                //— единица измерения (l, ml, g, kg, cnt или tens);
		pri:     stringToFloat(s2[3]), // содержание белков, шесть знаков после запятой
		fi:      stringToFloat(s2[4]), // содержание жиров
		chi:     stringToFloat(s2[5]), // содержание углеводов
		fvi:     stringToFloat(s2[6]),
	}
	return c
}
func countingCalories(i ingredient, c calories) caloriesDish {
	ml := c.ai
	var cd caloriesDish

	fmt.Printf("сравниваем %v и %v \n", i.nameIng, c.nameIng)
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
	fmt.Println(cd)

	//надо посчитать бжу 50 гр сыр
	//на 1 кг сыр б 20 ж 30 у 4
	//на 50 гр 20/100*50
	return cd
}
func sumCalories(s caloriesDish, s2 caloriesDish) caloriesDish {
	s.chi = s.chi + s2.chi
	s.fi = s.fi + s2.fi
	s.pri = s.pri + s2.pri
	s.fvi = s.fvi + s2.fvi
	return s
}
