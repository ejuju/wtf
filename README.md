
Roadmap:
- [ ] Write unit tests (WIP)
- [ ] Sine function (wavy GIFs)
- [ ] Combining two pictures with noise
- [ ] Save configuration and other to file when generating art (to be able to reproduce it) maybe write os.Stdout output to a file (and println config)
- [ ] Use goroutines to improve speed (see if maybe we can more space => dynamic progr. to save time, there may be some duplicate calculations in the pixel modifier logic)
- [ ] Set defaults to funcs starting New***
- [ ] Compare speed of encoding jpeg or png
- [ ] Make sure noise generators generate values within the same range (for ex: between -1 and 1)