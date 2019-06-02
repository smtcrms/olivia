package training

import (
	"github.com/gookit/color"
	"github.com/olivia-ai/gonn/gonn"
	"github.com/olivia-ai/olivia/analysis"
	"github.com/olivia-ai/olivia/util"
	"os"
)

// Return the inputs and targets generated from the intents for the neural network
func TrainData() (inputs, targets [][]float64) {
	words, classes, documents := analysis.Organize()

	for _, document := range documents {
		outputRow := make([]float64, len(classes))
		bag := document.Sentence.WordsBag(words)

		// Change value to 1 where there is the document Tag
		outputRow[util.Index(classes, document.Tag)] = 1

		// Append data to trainx and trainy
		inputs = append(inputs, bag)
		targets = append(targets, outputRow)
	}

	return inputs, targets
}

// Returns a new neural network and learn from the TrainData()'s inputs and targets
func CreateNeuralNetwork() (network gonn.NeuralNetwork) {
	// Decide if the network is created by the save or is a new one
	saveFile := "res/training.json"

	_, err := os.Open(saveFile)
	if err != nil {
		// Train the model if there is no training file
		trainx, trainy := TrainData()
		inputLayers, outputLayers := len(trainx[0]), len(trainy[0])
		hiddenLayers := 100

		network = *gonn.DefaultNetwork(inputLayers, hiddenLayers, outputLayers, true)
		network.Train(trainx, trainy, 1000)
		gonn.DumpNN(saveFile, &network)
	} else {
		color.FgBlue.Println("Loading the neural network from res/training.json")
		network = *gonn.LoadNN(saveFile)
	}

	return network
}
