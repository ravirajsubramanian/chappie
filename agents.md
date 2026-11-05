# Agent Context: Chappie CLI

## Overview
The Chappie CLI is a simple Go-based command-line utility designed for scheduled, repetitive tasks. Its primary function is to act as a simple, customizable job scheduler, running a specific action once per hour.

## Core Design Principles
1.  **Modularity:** The main execution loop is separate from the business logic (`runAction()` in `main.go`), allowing future functions (e.g., file processing, API calls, etc.) to be easily plugged in without altering the scheduling mechanism.
2.  **Simplicity:** The current implementation focuses on the basic structure (setup, logging, looping) rather than complex task execution.
3.  **Extensibility:** The `runAction()` function is the designated point of modification for all future features.

## Setup and Run Instructions (For reference in documentation)
See the `README.md` for detailed setup and usage instructions. The current implementation greets the user with "hello" and executes this greeting on startup and every hour until midnight.

---
*Last updated: Nov 5, 2026*