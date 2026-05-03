package main

import (
	"syscall"
	"unsafe"
)

type tWidth struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getTerminalWidth() int {
	var width tWidth

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(1), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&width)))
	if err != 0 {
		return 80
	}
	return int(width.Row)

}

/*

Your Go Program
      ↓
   &width          → find where width lives in memory
      ↓
unsafe.Pointer     → make it a generic address
      ↓
  uintptr          → make it a raw number
      ↓
syscall.Syscall    → hand everything to the OS
      ↓
   The OS          → reads terminal, writes into &width
      ↓
  width.Col        → your answer is here
      ↓
  return int       → give it back as a usable Go integer


	syscall.Syscall(
    syscall.SYS_IOCTL,              // "I want to control a device"
    uintptr(1),                      // "the device is my terminal"
    syscall.TIOCGWINSZ,             // "specifically get its window size"
    uintptr(unsafe.Pointer(&width)) // "write the answer here"
)

What Happens Step by Step

Your program calls syscall.Syscall
Your program pauses
OS takes over, finds your terminal
OS reads the terminal dimensions
OS writes 8 bytes into the address you provided (&width)
OS hands control back to your program
Your program reads width.Col — the OS already filled it in


Concept 4: The Operating System and System Calls

What is the OS?
Your computer has two worlds:
User Space — where your program lives. Go code, Python scripts, apps. Everything you write runs here.
Kernel Space — where the OS lives. It controls hardware, memory, files, network, terminal. Your program cannot directly touch anything here.
Think of it like a bank:

You are a customer (your program)
The vault is hardware/OS resources
You cannot walk into the vault yourself
You must ask a bank teller to do things for you


What is a System Call?
A system call is your program formally requesting the OS to do something it cannot do itself. Like filling out a bank request form and handing it to the teller.
Every time your Go program does something like:

Read a file → system call
Write to terminal → system call
Get terminal size → system call

Your program stops, hands control to the OS, the OS does the work, then hands control back to your program with the result.

What is a File Descriptor?
The OS manages every open resource with a numbered ticket:

0 → stdin (keyboard)
1 → stdout (terminal output)
2 → stderr (error output)
3, 4, 5... → files you open, network connections etc.

When you pass uintptr(1) to your syscall you're saying "I'm talking about stdout — my terminal".

What is IOCTL?
ioctl stands for Input Output ConTroL. It's a special system call that lets you query or control device settings. Think of it as a Swiss army knife system call — one tool for many device-related questions.
You tell it:

Which device — via file descriptor (1 = your terminal)
What you want to know — via a request constant (TIOCGWINSZ = get window size)
Where to put the answer — via your struct's address


What is TIOCGWINSZ?
It's a constant that breaks down as:

Terminal
IO
Control
Get
WINdow
SiZe

It's just a number that the OS recognizes as "the caller wants the terminal dimensions". Without it, the OS wouldn't know what device information you're asking for.
*/
