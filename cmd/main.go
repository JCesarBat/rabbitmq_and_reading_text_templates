package main

import (
	"Practics_with_templates/internal/reader"
	"Practics_with_templates/internal/taskDistributor"
	"fmt"
	"sync"
)

// TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
const testTemplateFilepath = "internal/seed/test_template.tpl"
const AnotherTestTemplate = "internal/seed/another_test_template"

func main() {

	distributor, err := taskDistributor.NewTaskDistributor()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	go distributor.Start()

	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan error, 2)
	go func(ch chan error) {
		defer wg.Done()
		for {

			data := reader.DefaultTestTemplate()

			output, err := reader.Read(testTemplateFilepath, data)
			if err != nil {
				ch <- err
				return
			}
			err = distributor.SendQ(output, testTemplateFilepath)
			if err != nil {
				ch <- err
				return
			}
		}
	}(ch)
	go func(ch chan error) {
		for {

			defer wg.Done()
			output, err := reader.FuncReader(AnotherTestTemplate)
			if err != nil {
				ch <- err
				return
			}
			err = distributor.SendQ(output, AnotherTestTemplate)
			if err != nil {
				ch <- err
				return
			}
		}

	}(ch)

	go func(ch chan error) {

		for {
			err := <-ch
			println(err.Error())
			if _, ok := <-ch; !ok {
				println("the channel is closed")
				return
			}
		}
	}(ch)

	go distributor.Consumer(AnotherTestTemplate)
	go distributor.Consumer(testTemplateFilepath)
	wg.Wait()
	close(ch)

	distributor.Done <- true

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
