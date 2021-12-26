package lib

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func ProcessingStrTolist2dComplex(text []byte) [][]complex128 {
	list2dComplex := new([][]complex128)
	listComplex := new([]complex128)
	list_ar_complex_in_str := strings.Split(string(text), ";")
	_, errorr := strconv.ParseComplex("end\n", 128)
	for _, value := range list_ar_complex_in_str {
		tmpl_rune := []rune(value)
		list_items := strings.Split(string(tmpl_rune), ",")
		for _, item := range list_items {
			temple, err := strconv.ParseComplex(item, 128)
			switch {
			case err == nil:
				*listComplex = append(*listComplex, temple)
			case reflect.DeepEqual(err, errorr):
				fmt.Println("Чтение из скрипта завершено")
			default:
				fmt.Println("Произошла ошибка при чтении ->", err)
			}

		}

		if len(*listComplex) != 0 {
			*list2dComplex = append(*list2dComplex, *listComplex)
			*listComplex = nil

		}

	}

	return *list2dComplex
}

func PocessingStr() []complex128 {
	/* Функция принимает на вход(поток ввода) строку из пайтон скрипта, в которой содержиться список чисел complex128,
	разделенных символом ";". Это стока пребразуется в []complex128
	*/
	fmt.Println("Запустился main.go")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(text)
	strComplexProgessing := strings.Split(string(runes), ";")

	listComplex := new([]complex128)

	for _, value := range strComplexProgessing {
		temple, err := strconv.ParseComplex(value, 128)
		if err == nil {
			*listComplex = append(*listComplex, temple)
		} else {
			fmt.Println("Чтение завершено ->", err)
		}

	}
	return *listComplex
}
