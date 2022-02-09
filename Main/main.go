package main

import (
	"FirstPrototip/Main/lib"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Box struct {
	frame      *fyne.Container
	fieldEntry []*widget.Entry
	fieldLabel []*widget.Label
}

type InfoAboutPolynom struct {
	PolynomialCoefficients string
	MultiplicityRoots      map[string]int
	AllRoots               string
}

func GlobaloAlphaModel(listComplexParametrs [][]complex128, psi float64, r float64, amounrSpoint,
	amountParticle, iter int, progress *widget.ProgressBar) []byte {
	listInfoPoly := new([]InfoAboutPolynom)
	h := 1 / float64(len(listComplexParametrs))
	i := 0.0
	for _, ar := range listComplexParametrs {
		roots := lib.ModelAlpha(ar, psi, r, amounrSpoint, amountParticle, iter)
		multiplicityRoots, flac := lib.MultiplicityOfRoots(ar, roots)
		elemet := InfoAboutPolynom{lib.ConvListComplexToString(ar), multiplicityRoots, flac}
		*listInfoPoly = append(*listInfoPoly, elemet)
		i += h
		progress.SetValue(i)

	}

	data, _ := json.MarshalIndent(listInfoPoly, "", "  ")
	//if err := ioutil.WriteFile("./json_files/list_info_polynom.txt", data, 0600); err != nil {
	//}
	return data

}

func eventToPushButtonStart(progress *widget.ProgressBar, dict map[string][]float64, listComplexParametrs [][]complex128) []byte {
	data := GlobaloAlphaModel(listComplexParametrs, dict["psi"][0], dict["r"][0], int(dict["point"][0]),
		int(dict["particle"][0]), int(dict["iter"][0]), progress)

	return data

}

func readParam(box Box) (map[string][]float64, bool, *[]float64) {
	flag := false
	endParam := new(float64)
	amountList := new(uint8)
	dict := make(map[string][]float64)
	variable := new([]float64)

	for ind, entryLocal := range box.fieldEntry {
		floatNum, er := strconv.ParseFloat(entryLocal.Text, 64)
		if er != nil {
			listStr := strings.Split(entryLocal.Text, ",")
			if len(listStr) != 2 {
				break
			} else {
				runef := []rune(listStr[0])
				runes := []rune(listStr[1])
				stParam, er_1 := strconv.ParseFloat(strings.ReplaceAll(string(runef), "[", ""), 64)
				enParam, er_2 := strconv.ParseFloat(strings.ReplaceAll(string(runes), "]", ""), 64)
				if er_1 == nil && er_2 == nil && stParam <= enParam {
					*amountList++
					endParam = &enParam
					*variable = []float64{stParam, enParam}
					dict[box.fieldLabel[ind].Text] = *variable

				} else {
					break
				}
			}

		} else {
			dict[box.fieldLabel[ind].Text] = []float64{floatNum}
		}
	}

	if *amountList == 1 && len(box.fieldLabel) == len(dict) && len(dict["h"]) == 1 && dict["h"][0] <= *endParam {

		flag = true

	}

	return dict, flag, variable
}

func createObjectMenuFromButton(mainWindow, localWindow fyne.Window, app fyne.App, box Box, name string) *widget.Button {
	buttonBack := widget.NewButton("back to menu", func() {
		localWindow.Hide()
		mainWindow.Show()
	})

	buttonForModel := widget.NewButton(name, func() {
		mainWindow.Hide()
		localWindow.Resize(fyne.NewSize(300, 450))

		label := widget.NewLabel("")

		progress := widget.NewProgressBar()

		button := widget.NewButton("start", func() {

			//fmt.Println(box.fieldLabel)

			dict, flag, variable := readParam(box)
			//fmt.Println(dict)

			if flag {
				label.SetText("correct")
				pack := make(map[string][]float64)
				pack["-v"] = *variable
				pack["-H"] = dict["h"]

				jsonDict, err := json.Marshal(dict) //Indent(dict, "", "\t")
				if err != nil {
					log.Fatal(err)
				}
				packJsonDict, errr := json.Marshal(pack) //Indent(dict, "", "\t")
				if err != nil {
					log.Fatal(errr)
				}

				//fmt.Println(string(packJsonDict))
				label.SetText("calculating coefficents")
				cmd, _ := exec.Command("python3", "Main/calcMass.py", "-im", string(jsonDict)).Output()
				//fmt.Println("done")
				//fmt.Println(string(cmd))

				listComplexParametrs := lib.ProcessingStrTolist2dComplex(cmd)
				//fmt.Println(listComplexParametrs)

				data := eventToPushButtonStart(progress, dict, listComplexParametrs)

				//data := lib.GlobaloAlphaModel(listComplexParametrs, dict["psi"][0], dict["r"][0], int(dict["point"][0]),
				//	int(dict["particle"][0]), int(dict["iter"][0]))

				//fmt.Println(string(data))
				label.SetText("waiting create pictures")
				cmd2, er := exec.Command("python3", "Main/Painter.py", "-j", string(data), "-p", string(packJsonDict)).Output()
				if er == nil {
					fmt.Println(string(cmd2))
					fmt.Println("end program")
					label.SetText("")
					progress.SetValue(0)
				} else {
					fmt.Println(er)
					fmt.Print(" wrong on python script")
					label.SetText("render time error")
				}

			} else {
				label.SetText("write to only corect types!")

			}
		})

		contentList := container.NewVBox(
			box.frame,
			label,
			progress,
			button,
			buttonBack,
			/*
				container.NewHBox(
					button,
					buttonBack),
			*/

		)

		grid := container.New(layout.NewCenterLayout(), contentList)

		localWindow.SetContent(grid)

		localWindow.SetCloseIntercept(func() {
			app.Quit()
		})

		localWindow.Show()
	})

	return buttonForModel
}

func main() {
	/*
		fmt.Println("lauching main.go")
		cmd, _ := exec.Command("python3", "Main/genMass.py").Output()
		fmt.Println(string(cmd))
		listComplexParametrs := lib.ProcessingStrTolist2dComplex(cmd)
		lib.GlobaloAlphaModel(listComplexParametrs, 1e-10, 10, 3000, 16, 60)
		//fmt.Println(string(data))
		cmd2, er := exec.Command("python3", "Main/Painter.py").Output()
		if er == nil {
			//fmt.Println("done")
			fmt.Println(string(cmd2))
		} else {
			fmt.Println("wrong!")
			fmt.Println(er)

		}
	*/
	application := app.New()
	mainWindow := application.NewWindow("Menu")
	mainWindow.Resize(fyne.NewSize(140, 100))

	label1_1 := widget.NewLabel("alpha")
	entry1_1 := widget.NewEntry()
	entry1_1.SetText("[1,10]")

	label1_2 := widget.NewLabel("A")
	entry1_2 := widget.NewEntry()
	entry1_2.SetText("1")

	label1_3 := widget.NewLabel("b_0x")
	entry1_3 := widget.NewEntry()
	entry1_3.SetText("1")

	label1_4 := widget.NewLabel("b_0y")
	entry1_4 := widget.NewEntry()
	entry1_4.SetText("1")

	label1_5 := widget.NewLabel("b")
	entry1_5 := widget.NewEntry()
	entry1_5.SetText("1")

	label1_6 := widget.NewLabel("k")
	entry1_6 := widget.NewEntry()
	entry1_6.SetText("1")

	label1_7 := widget.NewLabel("l")
	entry1_7 := widget.NewEntry()
	entry1_7.SetText("1")

	label1_8 := widget.NewLabel("H_0")
	entry1_8 := widget.NewEntry()
	entry1_8.SetText("1")

	label1_9 := widget.NewLabel("g")
	entry1_9 := widget.NewEntry()
	entry1_9.SetText("1")

	label1_10 := widget.NewLabel("Rm")
	entry1_10 := widget.NewEntry()
	entry1_10.SetText("1")

	label1_11 := widget.NewLabel("mu")
	entry1_11 := widget.NewEntry()
	entry1_11.SetText("1")

	label1_12 := widget.NewLabel("rho")
	entry1_12 := widget.NewEntry()
	entry1_12.SetText("1")

	label3_1 := widget.NewLabel("psi")
	entry3_1 := widget.NewEntry()
	entry3_1.SetText("1e-8")

	label3_2 := widget.NewLabel("r")
	entry3_2 := widget.NewEntry()
	entry3_2.SetText("5")

	label3_3 := widget.NewLabel("h")
	entry3_3 := widget.NewEntry()
	entry3_3.SetText("0.5")

	label4_1 := widget.NewLabel("iter")
	entry4_1 := widget.NewEntry()
	entry4_1.SetText("20")

	label4_2 := widget.NewLabel("particle")
	entry4_2 := widget.NewEntry()
	entry4_2.SetText("8")

	label4_3 := widget.NewLabel("point")
	entry4_3 := widget.NewEntry()
	entry4_3.SetText("500")

	label_entry := widget.NewLabel("")

	container1_l := container.NewVBox(label1_1, label1_2, label1_3, label1_4,
		label1_5, label1_11, label_entry,
		label3_1, label3_2, label3_3,
	)
	container1_e := container.NewVBox(entry1_1, entry1_2, entry1_3, entry1_4,
		entry1_5, entry1_11, label_entry,
		entry3_1, entry3_2, entry3_3,
	)

	container2_l := container.NewVBox(label1_6, label1_7, label1_8, label1_9,
		label1_10, label1_12, label_entry,
		label4_1, label4_2, label4_3,
	)
	container2_e := container.NewVBox(entry1_6, entry1_7, entry1_8, entry1_9,
		entry1_10, entry1_12, label_entry,
		entry4_1, entry4_2, entry4_3,
	)

	//container3_l := container.NewVBox(label3_1, label3_2, label3_3)
	//container3_e := container.NewVBox(entry3_1, entry3_2, entry3_3)

	//container4_l := container.NewVBox(label4_1, label4_2, label4_3)
	//container4_e := container.NewVBox(entry4_1, entry4_2, entry4_3)

	grid1 := container.New(layout.NewGridLayout(4), container1_l, container1_e, container2_l, container2_e)
	//	container3_l, container3_e, container4_l, container4_e)

	Tbox1 := Box{grid1, []*widget.Entry{entry1_1, entry1_2, entry1_3,
		entry1_4, entry1_5, entry1_6, entry1_7, entry1_8, entry1_9, entry1_10, entry1_11, entry1_12,
		entry3_1, entry3_2, entry3_3, entry4_1, entry4_2, entry4_3,
	},

		[]*widget.Label{label1_1, label1_2, label1_3,
			label1_4, label1_5, label1_6, label1_7, label1_8, label1_9, label1_10, label1_11, label1_12,
			label3_1, label3_2, label3_3, label4_1, label4_2, label4_3,
		}}
	firstmodelWindow := fyne.CurrentApp().NewWindow("First model")
	buttonFirstModel := createObjectMenuFromButton(mainWindow, firstmodelWindow, application, Tbox1, "btn1")

	/*
		label2_1 := widget.NewLabel("param1")
		entry2_1 := widget.NewEntry()

		label2_2 := widget.NewLabel("param2")
		entry2_2 := widget.NewEntry()

		label2_3 := widget.NewLabel("param3")
		entry2_3 := widget.NewEntry()

		label2_4 := widget.NewLabel("param4")
		entry2_4 := widget.NewEntry()

		container2_l := container.NewVBox(label2_1, label2_2, label2_3, label2_4)
		container2_e := container.NewVBox(entry2_1, entry2_2, entry2_3, entry2_4)

		grid2 := container.New(layout.NewGridLayout(2), container2_l, container2_e)
		secondModelWindow := fyne.CurrentApp().NewWindow("Second model")
		buttonSecondModel := createObjectMenuFromButton(mainWindow, secondModelWindow, application, grid2, "btn2")

		label3_1 := widget.NewLabel("param1")
		entry3_1 := widget.NewEntry()

		label3_2 := widget.NewLabel("param2")
		entry3_2 := widget.NewEntry()

		label3_3 := widget.NewLabel("param3")
		entry3_3 := widget.NewEntry()

		label3_4 := widget.NewLabel("param4")
		entry3_4 := widget.NewEntry()

		container3_l := container.NewVBox(label3_1, label3_2, label3_3, label3_4)
		container3_e := container.NewVBox(entry3_1, entry3_2, entry3_3, entry3_4)

		grid3 := container.New(layout.NewGridLayout(2), container3_l, container3_e)
		thirdModelWindow := fyne.CurrentApp().NewWindow("Third model")
		buttonThirdModel := createObjectMenuFromButton(mainWindow, thirdModelWindow, application, grid3, "btn3")
		buttonThirdModel.Resize(fyne.NewSize(190, 300))


	*/

	/*
		itemList := container.NewVBox(
			buttonFirstModel,
			buttonSecondModel,
			buttonThirdModel,
		)

		Menu := container.New(layout.NewCenterLayout(), itemList)
		mainWindow.SetContent(Menu)
	*/

	mainWindow.SetContent(container.NewVBox(
		buttonFirstModel,
		//buttonSecondModel,
		//buttonThirdModel,
	))

	mainWindow.SetCloseIntercept(func() {
		application.Quit()
	})

	mainWindow.ShowAndRun()
}
