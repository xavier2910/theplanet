package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf(`
 There is a roar of wind as the force of atmospheric entry robs you of your weightlessness.
Your descent pod rattles and shakes with the deceleration. Out of the pod's one window you
can see the glowing flames of your descent. Plummeting through the atmosphere, the pod's
descent is abruptly halted by the snap of the parachute opening. The pod falls gently, landing
with a mild bump. Through the window you can see the bright sky and stark purple rocks of

The Planet

A game by Xavier

Version: as yet unpublished.

[press enter]
`)

	fmt.Scanln()

	fmt.Printf(`
 As it turns out, some important people in charge of the planetary defense of the earth got
wind of your exploits in deep space. It turns out that these yellow furry aliens are rather a
menace to society and need to be eliminated. As the only one to have encountered them and
come back alive, you were considered an expert in the field and promptly dispatched to the
purple whale planet, which your teleportation from was tracked to. Thus,
`)

	start()
}

type p struct {
	T string
	G func()
}

func dispatch(paths ...p) {

	for i, elem := range paths {
		fmt.Printf("\t--%d: %s\n", i+1, elem.T)
	}

	gotGoodInput := false
	var input int
	for !gotGoodInput {
		fmt.Print("> ")
		fmt.Scanf("%d\n", &input)
		gotGoodInput = input > 0 && input <= len(paths)
	}
	paths[input-1].G()
}

func jumpTo(g func()) {

	fmt.Printf("\n[press enter]")
	fmt.Scanln()
	g()
}

func playAgainPrompt() {

	fmt.Printf("\nPlay again?\n")
	dispatch(
		p{"Yes", start},
		p{"No", func() {
			os.Exit(0)
		}},
	)
}

func start() {

	fmt.Printf(`
 You are sitting in a space pod with walls so thick you can only see out through the window.
Fortunately, it is well lit, and your hand is quite visible in front of your face. Unlike
some space pods you could mention, all its functionality is well labeled, and it is
guaranteed not to disappear while you're not looking. Out the porthole you can see a barren
purple landscape.
[Type the number of the option you'd like to choose, then press enter]
`)
	dispatch(
		p{T: "Exit pod", G: outsidePod},
		p{T: "Launch into space!!", G: launchToSpace},
		p{T: "Have an existential argument with the onboard computer", G: argueWithCptr},
	)
}

func argueWithCptr() {

	fmt.Printf(`
 The argument is prolix and protracted. Nothing much is gained.
`)

	jumpTo(start)
}

func launchToSpace() {

	fmt.Printf(`
 The rocket engines on your pod roar as you blast into space, leaving the purple landscape
behind. Soon you are in orbit.
`)
	jumpTo(spaceLow)
}

func spaceLow() {

	fmt.Printf(`
 Your pod hangs low in space over the purple rock of a planet beneath you. The navigation
terminal is a little simple-minded; it's monitor only shows:
`)
	dispatch(
		p{"Land on the purple planet", deltaVCrisis},
		p{"Return to earth", getCourtMartialed},
	)
}

func deltaVCrisis() {

	fmt.Printf(`
 The flight computer revolts against such a callous waste of Δv and depressurizes the cabin. You
die.
	The end.
`)
	playAgainPrompt()
}

func getCourtMartialed() {

	fmt.Printf(`
 You return to earth, but the powers that be are not happy with having spent billions of dollars to
send you to space only to have you return, having done nothing. You are executed. You
die.

	The end.
`)
	playAgainPrompt()
}

func outsidePod() {

	fmt.Printf(`
 You are standing just outside your space pod, in the midst of a barren purple plain. Bluegrass shivers
in the wind, which is apparently breathable, since you aren't dead. Off to your left are some hills; dead
ahead, some large beasts are grazing, green in color; apparently they have no need for camouflage.
`)
	dispatch(
		p{"Return to pod", start},
		p{"Go to hills", hills},
		p{"Go to the large beasts", beasts},
		p{"Run around in circles", runInCircles},
	)
}

func runInCircles() {

	fmt.Printf(`
 You run around in circles, waving your arms and screaming, just for kicks. Nothing happens.
`)
	jumpTo(outsidePod)
}

func hills() {

	fmt.Printf(`
 You are standing atop a rocky purple hill in the midst of a bluegrass plain. From this high
vantage point you can see your landed pod, a herd of great green beasts grazing on the grass,
and a town or settlement of some kind off in the distance, though with the purple haze it's
hard to make out just how big it is.
`)
	dispatch(
		p{"Return to pod", outsidePod},
		p{"Get a closer look at the beasts", beasts},
		p{"Go off toward the settlement", town},
	)
}

func beasts() {
	fmt.Printf(`
 As you approach the beasts, some of the aliens, wearing blue all over, and riding some kind of
large blue fish, come galloping in and shoot one of the beasts, causing the rest to stampede.
Fortunately, away from you. Unfortunately, they quickly spot you in your orange flight suit. One
of them points his weapon at you.
`)
	dispatch(
		p{"Beg for mercy!", mercyFromCowboys},
		p{"Shoot them first!!", shootCowboys},
		p{"Run!!!", runFromCowboys},
	)
}

func mercyFromCowboys() {

	fmt.Printf(`
 They appear unable to understand your requests. They tie you up and carry you off to the settlement.
`)
	jumpTo(trial)
}

func runFromCowboys() {
	fmt.Printf(`
 You run as fast as you can, but their fish are still faster. They tie you up and carry your off to
the settlement.
`)
	jumpTo(trial)
}

func shootCowboys() {
	fmt.Printf(`
 Unlike some previous ventures, you are actually armed this time! Oops, that one's a faster draw. You
die.

	The end.
`)
	playAgainPrompt()
}

func trial() {
	fmt.Printf(`
 You are transferred to a cage in an open square. Soon, a veritable army of yellow furry aliens come
marching out of various buildings, all in impressive looking green suits, evidently made from the fur
of creatures like the ones you saw on the plains. A crowd of spectators gathers; they are wearing
various colors, but very few wear the green that characterizes the impressive agents of the alien
judicial apparatus. They conduct what appears to be a trial or other public discussion as to your fate.
In the end, they whack you several times with something hard and let you go.
`)
	jumpTo(town)
}

func town() {
	fmt.Printf(`
 You are standing at the outskirts of a small town thingy. Various yellow-haired furry aliens go
about their business in the streets, walking hither and yon. Behind you is a lonely group of purple
rocky hills.
`)
	dispatch(
		p{"Enter town", townSquare},
		p{"Go around town", aroundTown},
		p{"Go to hills", hills},
	)
}

func aroundTown() {
	fmt.Printf(`
 You skirt the edges of the town, and eventually come to what looks like a spaceport.
`)
	jumpTo(spaceport)
}

func townSquare() {
	fmt.Printf(`
 You are in what appears to be the center of the town, with several buildings facing you. One tall
impressive building with green lettering (that you cannot read) looms above the rest. A pathway
leads back out to the bluegrass prairie, and another further on into the town.
`)
	dispatch(
		p{"Enter impressive building", lobby},
		p{"Go to the edge of town", town},
		p{"Go further on", spaceport},
	)
}

func lobby() {
	fmt.Printf(`
 You enter the building, and come to a grand lobby with a desk. At the desk is a relatively small
yellow-furred figure, eyeing you warily.
`)
	dispatch(
		p{"Greet the clerk", greetClerk},
		p{"Slowly back away and hope he doesn't notice you", townSquare},
	)
}

func greetClerk() {
	fmt.Printf(`
 The yellow alien looks up at you in shock. For several seconds.
`)

	jumpTo(clerkTalks)
}

func clerkTalks() {
	fmt.Printf(`
 "Wha--! Hm. Name? Occupation? Planet of origin? Business with the central government of
all Gorbluks?" he says, clearly trying to wrap his head around the whole situation.
`)
	dispatch(
		p{"\"Bob, Ambassador, Earth, to advocate for the same before the government.\"", stateBusiness},
		p{"\"Why, in the name of all the emptiness of the void, are you the first" +
			"\n\t    one around here who listens to me!?\"", questionComms},
		p{"Turn around and leave", townSquare},
	)
}

func questionComms() {
	fmt.Printf(`
 "You certainly don't *look* like someone who could engage in intelligent conversation! How
was anyone to know!? Now state your business or leave!"
`)
	dispatch(
		p{"\"Bob, Ambassador, Earth, to advocate for the same before the government.\"", stateBusiness},
		p{"Turn around and leave", townSquare},
	)
}

func stateBusiness() {
	fmt.Printf(`
 "Oh? Very well," says the Gorbluk clerk with an odd expression on his face. "Follow
me."

 He leads you through some doors and up (and down) some stairs, and plops you down on a
chair in a small, round room. He then leaves, shutting the door behind him. Suddenly, the
walls open up, and a booming voice declares,

 "YOU ARE NOW IN THE PRESENCE OF HIS MIGHTY MAJESTY, THE LORD HIGH DICTATOR OF ALL GORBLUKS!
KNEEL, PUNY CREATURE, OR BE MADE NON-EXISTENT!!"
`)
	// commercial break ;)
	jumpTo(thePresence)
}

func spaceport() {
	fmt.Printf(`
 You are standing in the midst of what seems to be a spaceport, with hangars and launch pads
scattered hither and yon. A relatively few aliens are going about servicing the spacecraft.
`)
	dispatch(
		p{"Commandeer that one standing ready to launch and fly away!", spacePiracy},
		p{"Go into the town", townSquare},
		p{"Go around the town", town},
	)
}

func spacePiracy() {
	fmt.Printf(`
 You manage to sneak on to one of the spacecraft. In it, you are surrounded by unfamiliar gauges
and instruments. But you can pick out three vaguely recognizable instruments: a big yellow button,
a small green button, and a red lever marked emergency.
`)
	dispatch(
		p{"Push the yellow button!", launchAlienShip},
		p{"Push the green button!", exitToSpaceport},
		p{"Pull the red lever!", ejectAtSpaceport},
	)
}

func ejectAtSpaceport() {
	fmt.Printf(`
 With a terrific bang, the wall flies off of the side of the spacecraft and you are thrown bodily
out of the craft, falling to the ground, unconscious.
`)

	jumpTo(executed)
}

func exitToSpaceport() {
	fmt.Printf(`
 The green button opens the hatch. You try to sneak back out, but are caught by the authorities,
who knock you unconscious.
`)
	jumpTo(executed)
}

func executed() {
	fmt.Printf(`
 You come to yourself, tied up in a white room. A yellow-furred alien in a white coat bends over you,
and asks you how you feel about decapitation.
`)
	dispatch(
		p{"\"Wonderful\"", executed_Wonderful},
		p{"\"Disgusting\"", executed_Disgusting},
		p{"Refuse to reply", executed_Silent},
	)
}

func executed_Wonderful() {
	fmt.Printf(`
 "Ah, excellent!" says the alien, and decapitates you. You
die.

	The end.
`)
	playAgainPrompt()
}

func executed_Disgusting() {
	fmt.Printf(`
 "Ah, how ironic," says the alien, and decapitates you. You
die.

	The end.
`)
	playAgainPrompt()
}

func executed_Silent() {
	fmt.Printf(`
 "Huh. Not a talker," says the alien, and decapitates you. You
die.

	The (psych!) end.
`)
	playAgainPrompt()
}

func launchAlienShip() {
	fmt.Printf(`
 Apparently, the yellow button launches the alien spacecraft, lobbing you into deep
space. While the view out the window is quite spectacular, the sheer emptiness of space
makes it very unlikely that you will ever run into anything. You begin to get hungry.
`)
	dispatch(
		p{"Root around for food", getIntercepted},
		p{"Shout for help", getIntercepted},
		p{"Just sit there and contemplate your folly", getIntercepted},
	)
}

func getIntercepted() {
	fmt.Printf(`
 As you do, the onboard computer announces,

 "Ship approaching, enforcer class."

 Whereupon an enforcer craft docks with your spacecraft; aliens come aboard and tie you up
and knock you out.
`)
	jumpTo(executed)
}

func thePresence() {
	fmt.Printf(`
 You are standing in the presence of His Mighty Majesty, The Lord High Dictator Of All Gorbluks.
`)

	dispatch(
		p{"Kneel", thePresenceKneeling},
		p{"Be made non-existent", thePresenceObliterated},
	)
}

func thePresenceObliterated() {
	fmt.Printf(`
 You refuse to kneel, that being unfitting for an ambassador of your degree.

You are removed from the universe. The only way to return is to reboot.
`)

	os.Exit(0)
}

func thePresenceKneeling() {
	fmt.Printf(`
 You are kneeling in the presence of His Mighty Majesty, The Lord High Dictator Of All Gorbluks,
who condescends to address you thus:

 "Who are you, who thus comes to our Great Planet Of Glock, and Demands an Audience with Us And
Our Majesty!?"
`)
	dispatch(
		p{"\"O Mighty Dictator And Ruler Of The Universe, have" +
			"\n\t    Mercy upon the lowly planet of earth!\"", begForMercy},
		p{"Question his capitalization", questionCapitals},
	)
}

func questionCapitals() {
	fmt.Printf(`
 "You Know, I never Thought much about It. It is simply just how I speak. I think it seemed impressive
to Me when I became Dictator and all This was being Set Up. I Like Folks like You, who Think about
Stuff. So few people Think anymore, especially in My Presence. I Suspect, however, that some folks
Think when I'm not Looking. I Can't Understand it. In Any Case, I'm Sure more than My Grammar
brought You to Me. What is It, You Wanted Of Me?"
`)
	dispatch(
		p{"\"My Planet is Oppressed by Gorbluks. Could You please Stop them?\"", askForMercy},
		p{"\"Nothing, Sire and Lord; I simply found Your Grammar so Fascinating!\"", justGrammar},
	)
}

func justGrammar() {
	fmt.Printf(`
 "Oh! How Flattering! Remain in My Presence; I will pay you well."

 You live the rest of your life on the purple planet Glock as a jester to the Dictator. You are cared for
and given spending money. Earth is probably half-destroyed from Gorbluk menace, and a war has probably
started over there. You don't much care. You have a good job. Until you miscapitalize a word and are
executed. You
die.

	The end.
`)

	playAgainPrompt()
}

func askForMercy() {
	fmt.Printf(`
 "Your Petition is Granted. I will Withdraw All Forces from that region. And would You like to Return
Home to see My Promise Granted?"
`)
	dispatch(
		p{"\"Yes, O Most Mighty Majesty!\"", getSentHome},
		p{"\"No, grant Me rather to Remain in Your Presence for the Rest Of My Life," +
			"\n\t    for I have come to Appreciate Your Capitalization!\"", justGrammar},
	)
}

func getSentHome() {
	fmt.Printf(`
 "Ah, More's The Pity. Oh Well."

 The wall shuts and you are led to your waiting spacecraft, which takes you home to earth. You give
your report, but not being a close acquaintance of the Powers That Be in that part of the universe,
you return to your normal life.

	The end.
`)
	playAgainPrompt()
}

func begForMercy() {
	fmt.Printf(`
 "Why should I Show Mercy on a planet so irrelevant? Begone!"

 The wall shuts and you are led away, out of the impressive building, and into a land vehicle. The
land vehicle takes you a long, long distance. You are taken out of the vehicle, which then forsakes
you.
`)
	jumpTo(inDesert)
}

func inDesert() {
	fmt.Printf(`
 You are standing in the middle of a great, open expanse of pale purple desert. Around you rocks and
a bluish cactus adorn the barren sands. The hot wind stirs the sand up into the cloudless sky, in which
a whale flies by, singing its lonesome song, oblivious of all on the surface.
`)
	dispatch(
		p{"Hail the whale", hailWhale},
		p{"Shoot the whale", shootWhale},
		p{"Ignore the whale", ignoreWhale},
	)
}

func hailWhale() {
	fmt.Printf(`
 Surprisingly enough the whale does notice you calling out to it. It decides you are a new and
interesting food source and devours you. You
die.

	The (fishy) end.
`)
	playAgainPrompt()
}

func shootWhale() {
	fmt.Printf(`
 You draw your weapon and shoot at the whale. Though visibly injured, the whale is still able to go
about its normal business, such as eating you. You
die.

	The (fishy) end.
`)
	playAgainPrompt()
}

func ignoreWhale() {
	fmt.Printf(`
 In the absence of any indication of your presence, the whale swims away, leaving you alone with the
cactus and a rock.
`)
	dispatch(
		p{"Despair", dieOfDespair},
		p{"Attempt to harvest the cactus", func() { harvestCactus(false) }},
		p{"Look under the rock", func() { lookUnderRock(false) }},
	)
}

func dieOfDespair() {
	fmt.Printf(`
 You despair. You
die.

	The end.
`)
	playAgainPrompt()
}

func harvestCactus(lookedUnderRock bool) {
	fmt.Printf(`
 Although prickly, the cactus has some water in it, which you consume. Too bad you can only see one.
`)
	if lookedUnderRock {
		dispatch(
			p{"Despair", dieOfDespair},
		)
	} else {
		dispatch(
			p{"Despair", dieOfDespair},
			p{"Look under the rock", func() { lookUnderRock(true) }},
		)
	}
}

func lookUnderRock(harvestedCactus bool) {
	if harvestedCactus {
		fmt.Printf(`
 You lift the rock, and are suddenly thrown back. The world spins and your vision blurs, as you are
projected through space.
`)
		jumpTo(yellowPlanet)
	} else {
		fmt.Printf(`
 Under the rock is only sand.
`)
		dispatch(
			p{"Despair", dieOfDespair},
			p{"Attempt to harvest the cactus", func() { harvestCactus(true) }},
		)
	}
}

func yellowPlanet() {
	fmt.Printf(`
 You are standing in the midst of a great expanse of greenish-yellow. Bright yellow-orange grasses 
blow in the wind, which is also apparently breathable, though it tastes funny. Small pink animals
scurry through the underbrush. Blue-scaled aliens in yellow garb mounted on what look like miniature
gold-colored ostriches plunge through the grass, chasing the pink rodents, catching them on long
spears. On the horizon a huge battle-equipped spacecraft hovers amid the haze, painted a brilliant
cyan. The aliens to not appear to have noticed you.
`)
	dispatch(
		p{"Hail the aliens", hailKulbrogs},
		p{"Hide", hideFromKulbrogs},
		p{"Shoot the aliens", shootKulbrogs},
	)
}

func shootKulbrogs() {
	fmt.Printf(`
 You fire but miss. Alert, the aliens shoot you. You
die.

	The end.
`)
	playAgainPrompt()
}

func hideFromKulbrogs() {
	fmt.Printf(`
 You successfully hide from them continuously, but one of the pink rodents bites you and you catch a
fatal case of space rabies. Despite the best efforts of alien doctors, you
die.

	The end.
`)
	playAgainPrompt()
}

func hailKulbrogs() {
	fmt.Printf(`
 They ride over to you and dismount their strange steeds. One says to you,

 "Well met, extraplanetary stranger! We are Kulbrogs, inhabitants of this great planet Kcolg, and
members of its commonwealth. If you are friendly, we are willing to let you share our commonweal."

 You notice that, while all of them have spears, only one has a gun.
`)

	dispatch(
		p{"Yell, \"I reject your commonwealth!\" and shoot them.", shootKulbrogs},
		p{"\"I accept your offer as a representative of the planet Earth.\"", diplo},
		p{"\"What's with the names? They're just reversed of the Gorbluks'!\"", reversed},
	)
}

func reversed() {
	fmt.Printf(`
 "Ach! You've had a run in with those villains!? We hate them totally! We are dedicated to their
destruction! But intelligence never can find out where we ought to strike with our severely limited
firepower. If you've had a run in with them, do you have any information on them we could use?"
`)
	dispatch(
		p{"\"No.\"", rejection},
		p{"\"Why should I tell you? I *like* the Gorbluks!\"", stupid},
		p{"\"If we can find my old ship, I could navigate you to their capitol for" +
			"\n\t    its destruction!\"", offerHelp},
	)
}

func diplo() {
	fmt.Printf(`
 "Earth? That's rather an obscure place! I seem to remember there being some Gorbluk activity
there, eh Rigger?"

 One of the aliens nods.

 "Right-o, then. We have a common enemy, how 'bout we join forces, whaddaya say?"
`)
	dispatch(
		p{"\"If we can find my old ship, I could navigate you to their capitol for" +
			"\n\t    its destruction!\"", offerHelp},
		p{"\"I will remain neutral, so no.\"", rejection},
		p{"\"No, actually, I'm on the Gorbluk side here.\"", stupid},
	)
}

func stupid() {
	fmt.Printf(`
 The Kulbrogs shoot you. You
die.

Why did you think that was a good idea?

	The end.
`)
	playAgainPrompt()
}

func rejection() {
	fmt.Printf(`
 "Oh. Well. I guess we shouldn't have told you all we did. I'm sorry, but we're gonna have to kill
you now."

 You
die.

	The end.
`)
	playAgainPrompt()
}

func offerHelp() {
	fmt.Printf(`
 "Wonderful! Come with us."

 They lead you off across the plain towards the hovering battleship. A ladder comes down and you
are led up it and into the craft's living quarters. After a little while the ship begins to move.
Out your porthole you can see the ground drop away and be replaced by stars. After several weeks of
space travel, you finally arrive at Glock. You are brought to the bridge to guide the ship. The
spaceship bombards the central government of all Gorbluks. His Mighty Majesty, The Lord High Dictator
Of All Gorbluks dies. All Gorbluk activities come to a halt. The Earth is freed. The Kulbrogs invade
and lay waste to Glock. Through your artifice, a treaty is signed between the Earth and the
Commonwealth of Kcolg. You return to Earth, a hero. Eventually, the Earth becomes a satellite state
of the Kulbrogs. But their government isn't so bad, so long as nobody tells anyone anything they
shouldn't.

	The end.
`)
	playAgainPrompt()
}
