package main

// libraries needed
import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Shopify/go-lua"
)

// D4 function
func D4() int {
	return rand.Intn(4) + 1
}

// D6 function
func D6() int {
	return rand.Intn(6) + 1
}

// D8 function
func D8() int {
	return rand.Intn(8) + 1
}

// D10 function
func D10() int {
	return rand.Intn(10) + 1
}

// D12 function
func D12() int {
	return rand.Intn(12) + 1
}

// D20 function
func D20() int {
	return rand.Intn(20) + 1
}

// main function
func main() {
	// random seed
	rand.Seed(time.Now().UnixNano())

	// D4 channel declaration
	D4Chan := make(chan int)
	D6Chan := make(chan int)
	D8Chan := make(chan int)
	D10Chan := make(chan int)
	D12Chan := make(chan int)
	D20Chan := make(chan int)

	// Load lua standard libraries and open up lua VM for mod support
	l := lua.NewState()
	lua.OpenLibraries(l)

	// tells user where the game is looking for mod file
	cwd, _ := os.Getwd()
	fmt.Println("Looking for Mod file in:", cwd)

	// declaring variable for Mod option
	var intModOption int
	fmt.Println("Do you want to use mods, Yes(1) or No(2):")
	_, err := fmt.Scanln(&intModOption)
	if err != nil || intModOption < 1 || intModOption > 2 {
		fmt.Println("Not An Option")
		os.Exit(0)
	}

	// only used when player wants to use mod
	var script []byte
	if intModOption == 1 {
		script, err = os.ReadFile("The_Crusade_Campaign_For_Iberia_Video_Game_Mod.lua")
		if err != nil {
			panic(err)
		}

		// Execute the Lua mod script now to load globals before accessing them
		if err := lua.DoString(l, string(script)); err != nil {
			panic(err)
		}

		// Mod Name and Notice
		l.PushGlobalTable()
		l.PushString("strNoticeAndName")
		l.RawGet(-2)
		strModNoticeAndName, ok := l.ToString(-1)
		if !ok {
			fmt.Println("Failed to get Lua string: strNoticeAndName")
		} else {
			fmt.Println(strModNoticeAndName)
		}
		l.Pop(2)
	}

	// calls threads or cores for dice rolls
	go func() { D4Chan <- D4() }()
	go func() { D6Chan <- D6() }()
	go func() { D8Chan <- D8() }()
	go func() { D10Chan <- D10() }()
	go func() { D12Chan <- D12() }()
	go func() { D20Chan <- D20() }()

	// receives dice results from thread or core through Dice channel
	d4Result := <-D4Chan
	d6Result := <-D6Chan
	d8Result := <-D8Chan
	d10Result := <-D10Chan
	d12Result := <-D12Chan
	d20Result := <-D20Chan

	// forever loop needed for game
	for {
		// game info
		fmt.Println("This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the license, or (at your option) any later version. This program is distributed in hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have recieved a copy of the GNU General Public License along with this program. If not, see https://www.gnu.org/licenses")
		fmt.Println("Information just about the stuff in this software not covered by the GNU General Public License version 3 (i.e.: the ruleset of this game):  This work is licensed under Attribution-ShareAlike 4.0 International")
		fmt.Println("Copyright (C) 2025 Daniel Hanrahan Tools and Games")
		fmt.Println("You start in the city of Leon")
		fmt.Println("What city do you want to conquer (use city coresponding number as your option): ")
		fmt.Println("WARNING: You may have control some of these cities, write them down.")
		fmt.Println("1. Almeria")
		fmt.Println("2. Pampalona")
		fmt.Println("3. Zargosa")
		fmt.Println("4. Porto")
		fmt.Println("5. Barcelona")
		fmt.Println("6. Toledo")
		fmt.Println("7. Coimbra")
		fmt.Println("8. Santarem")
		fmt.Println("9. Lisbon")
		fmt.Println("10. Faro")
		fmt.Println("11. Cordoba")
		fmt.Println("12. Seville")
		fmt.Println("13. Grenada")
		fmt.Println("14. Baleric Islands")

		// Mod Option string print
		l.PushGlobalTable()
		l.PushString("strNewLocation")
		l.RawGet(-2)
		strModLocation, ok := l.ToString(-1)
		if !ok {
			fmt.Println("Failed to get Lua string: strNewLocation")
		} else {
			fmt.Println(strModLocation)
		}
		l.Pop(2)

		// only used when player wants to use mod
		if intModOption == 1 {
			// Run mod script each loop iteration (optional if you want updates every loop)
			if err := lua.DoString(l, string(script)); err != nil {
				fmt.Println("Lua Error:", err)
			}
		}

		// declaring variable for conquest option
		var intConquestOption int
		_, err := fmt.Scanln(&intConquestOption)
		if err != nil {
			fmt.Println("Not An Option")
			os.Exit(0)
		}

		if intModOption == 2 {
			if intConquestOption < 1 || intConquestOption > 14 {
				fmt.Println("Not An Option")
				os.Exit(0)
			}
		} else {
			if intConquestOption == 15 {
				l.PushGlobalTable()
				l.PushString("intDice")
				l.RawGet(-2)
				intModDice, ok := l.ToInteger(-1)
				if !ok {
					fmt.Println("Failed to get Lua integer: intDice")
				}
				l.Pop(2)

				if intModDice == 2 {
					l.PushGlobalTable()
					l.PushString("strBadResult")
					l.RawGet(-2)
					strModBadResult, ok := l.ToString(-1)
					if !ok {
						fmt.Println("Failed to get Lua string: strBadResult")
					} else {
						fmt.Println(strModBadResult)
					}
					l.Pop(2)
					os.Exit(0)
				}
			}
		}

		// letting player know, the army is split and is using all possible routes to get to the city as part of its strategy
		fmt.Println("Your army is split and is using all possible routes to get to the city as part of its strategy.")

		// terrain traversal check
		if d12Result != 12 {
			fmt.Println("Traversed enemy rivers and the seas successfully.")
			if d10Result != 10 {
				fmt.Println("Traversed the deserts successfully.")
				if d8Result != 8 {
					fmt.Println("Traversed allied rivers successfully.")
					if d6Result != 6 {
						fmt.Println("Traversed the forests successfully.")
						if d4Result != 4 {
							fmt.Println("Traversed the mountains successfully.")
							if d20Result == 20 {
								fmt.Println("You have successfully Conquered the city.")
							} else {
								fmt.Println("Your army got lost in terrain, Game Over")
								os.Exit(0)
							}
						} else {
							fmt.Println("Your army got lost in terrain, Game Over")
							os.Exit(0)
						}
					} else {
						fmt.Println("Your army got lost in terrain, Game Over")
						os.Exit(0)
					}
				} else {
					fmt.Println("Your army got lost in terrain, Game Over")
					os.Exit(0)
				}
			} else {
				fmt.Println("Your army got lost in terrain, Game Over")
				os.Exit(0)
			}
		} else {
			fmt.Println("Your army got lost in terrain, Game Over")
			os.Exit(0)
		}

		// political crisis check
		if d12Result != 12 {
			fmt.Println("No New Government.")
			if d10Result != 10 {
				fmt.Println("No Civil War.")
				if d8Result != 8 {
					fmt.Println("No annexation succession crisis.")
					if d6Result != 6 {
						fmt.Println("No breakaway succession crisis.")
						if d4Result != 4 {
							fmt.Println("No marriage annexation.")
							if d20Result != 20 {
								fmt.Println("No political revolution breakaway.")
							} else {
								fmt.Println("Political revolution breakaway, Game Over")
								os.Exit(0)
							}
						} else {
							fmt.Println("Marriage annexation, Game Over")
							os.Exit(0)
						}
					} else {
						fmt.Println("Breakaway succession crisis, Game Over")
						os.Exit(0)
					}
				} else {
					fmt.Println("Annexation succession crisis, Game Over")
					os.Exit(0)
				}
			} else {
				fmt.Println("Civil War, Game Over")
				os.Exit(0)
			}
		} else {
			fmt.Println("New Government, Game Over")
			os.Exit(0)
		}
	}
}
