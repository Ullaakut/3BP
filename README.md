# Three Body Problem

<p align="center"><img src="https://raw.githubusercontent.com/Ullaakut/3BP/master/3BP.gif" width="100%"/></p>

In physics and classical mechanics, the three-body problem is the problem of taking an initial set of data that specifies the positions, masses, and velocities of three bodies for some particular point in time and then determining the motions of the three bodies, in accordance with Newton's laws of motion and of universal gravitation which are the laws of classical mechanics. The three-body problem is a special case of the n-body problem. Unlike two-body problems, there is no general closed-form solution for every condition and numerical methods are needed to solve these problems.

## This is a work in progress

The goal of this project is to solve and render the orbits of three bodies in two dimensions in a terminal.

## How to test it?

* `get -u "github.com/nsf/termbox-go"`
* `git clone git@github.com:ullaakut/3BP`

## What is there left to do?

* [ ] Fix calculus. Currently using the parameters of the [Butterfly I configuration](http://three-body.ipb.ac.rs/sol.php?id=2) doesn't make the appropriate result. I think I got the equations incorrectly, and bodies are apparently not attracted to each other at the moment.
* [ ] Add interface to select configurations
* [ ] Add interface to manually change configuration
* [ ] Improve UI
