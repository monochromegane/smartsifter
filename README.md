# go-smartsifter

**SmartSifter: On-line Unsupervised Outlier Detection Using Finite Mixtures with Discounting Learning Algorithms.** This method is proposed by Yamanishi, K., Takeuchi, J., Williams, G. et al. (2004)

refs: [http://cs.fit.edu/~pkc/id/related/yamanishi-kdd00.pdf](http://cs.fit.edu/~pkc/id/related/yamanishi-kdd00.pdf)

![smartsifter](https://user-images.githubusercontent.com/1845486/36640826-93e5ab9a-1a69-11e8-8672-7b59116528ad.gif)

## Usage

```go
r := 0.1 // Discounting parameter.
alpha := 1.5 // Hyper parameter for continuous variables.
beta := 1.0 // Hyper parameter for categorical variables.
cellNum := 0 // Only continuous variables.
mixtureNum := 2 // Number of mixtures for GMM.
dim := 2 // Number of dimentions for GMM.

ss := smartsifter.NewSmartSifter(r, alpha, beta, cellNum, mixtureNum, dim)
logLoss := ss.Input(nil, []float64{0.1, 0.2}, true)
fmt.Println("Score using logLoss: %f\n", logLoss)
```

## Examples

Online outlier detection for faithful data (Only continuous variables and using parametric method).

```sh
$ go run examples/example.go
$ wget -O plot.py https://gist.github.com/monochromegane/8b6a2a18084297e05f3d25bde2518a9c
$ python plot.py
$ convert -adjoin out/*.jpg out.gif
```

If you want to try categorical variables, pass `-x` parameter to examples.go.


## TODO

- [ ] Implement non-parametric method (SDPU).
- [ ] Implement Hellinger score.
- [ ] Implement CLI.


## License

[MIT](https://github.com/monochromegane/smartsifter/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)
