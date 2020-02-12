//Nicholas Larsen
//February 11, 2020
//Is a Mad Lib and asks for specific terms to create a story


package main

import "fmt"

func main() {
var noun string
var number string
var verb string
var verb2 string
var relative string
//assigns all of the variables in the story

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my")

fmt.Println("Enter a noun of your choice")

fmt.Scanln(&noun)
//the first choice for the user

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the")

fmt.Println("enter a number of your choice")

fmt.Scan(&number)
//Second choice

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the",number,"th time.  I told the driver that he could go")

fmt.Println("enter a verb of your choice")

fmt.Scan(&verb)
//Third choice

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the",number,"th time.  I told the driver that he could go",verb,"himself and i was going to")

fmt.Println("enter another verb of your choice")

fmt.Scan(&verb2)
//Fourth choice

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the",number,"th time.  I told the driver that he could go",verb,"himself and i was going to",verb2,"his")

fmt.Println("enter a generic term for a relative; cousin, brother, etc.")

fmt.Scan(&relative)
//Last choice

fmt.Println("Yesterday, I got wicked pissed when a snowplow hit my",noun,"for the",number,"th time.  I told the driver that he could go",verb,"himself and i was going to",verb2,"his",relative,"Long story short, I need bail money, Love your Grandfather")
}