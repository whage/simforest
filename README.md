# Simforest: an ecological simulation

This is how I'm learning to program in Go.
Inspired by https://www.youtube.com/watch?v=r_It_X7v-1E. Go watch his videos!

## TODO
- implement needs
- make movement based on needs
- learn how to do "inheritance" in go
    - how to share code between related objects (as in parent class's method in classical OOP)
    - https://stackoverflow.com/questions/21251242/is-it-possible-to-call-overridden-method-from-parent-struct-in-golang
    https://stackoverflow.com/questions/37635769/elegant-way-to-implement-template-method-pattern-in-golang
    - https://github.com/bvwells/go-patterns
- use `float64` for Positions
- evolution of species
	- inheriting favorable abilities
- add diverse set of terrain types and creatures
- move to pixel-based graphics
	- raylib?
- implement an environment that is larger than the viewport (arbitrary world size)
	- make it possible to move the view
- procedurally generated terrain
- get rid of the inefficient searches (everything related to positions)
- implement a non-rectangular environment
	- give space to a HUD
