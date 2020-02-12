package main

import "fmt"

func main() {
var noun string
var number string
var verb string
var verb2 string
var relative string

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my")

fmt.Println("Enter a noun of your choice")

fmt.Scanln(&noun)

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the")

fmt.Println("enter a number of your choice")

fmt.Scan(&number)

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the",number,"th time.  I told the driver that he could go")

fmt.Println("enter a verb of your choice")

fmt.Scan(&verb)

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the",number,"th time.  I told the driver that he could go",verb,"himself and i was going to")

fmt.Println("enter another verb of your choice")

fmt.Scan(&verb2)

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the",number,"th time.  I told the driver that he could go",verb,"himself and i was going to",verb2,"his")

fmt.Println("enter a generic term for a relative; cousin, brother, etc.")

fmt.Scan(&relative)

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the",number,"th time.  I told the driver that he could go",verb,"himself and i was going to",verb2,"his",relative,"Long story short, I need bail money, Love your Grandfather")
}