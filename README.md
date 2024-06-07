# MonsterGenPro v0.1

## This project is a bit of a mess right now SORRY*
ps probs dont try and use it, part of it works but the API is BUSTED.

The application is designed to run a "server" that can randomly generate enemies or characters to populate books, games, or tabletop RPGs like D&D. Its core functionality revolves around the ability to upload spreadsheet data, which the program then stores and retrieves as needed. 

## Variations

One particularly intriguing feature is the "variation" mode, which allows the software to mix, match, and alter certain bits of information. This introduces more flexibility and randomness to the generated entities, making unique entities as needed!?

so what does this mean?

It's very simple. Lets say you have added a Goblin into your database, and you have vatiations turned on. 

The base Stat's are {'ID': 1, 'HP': 5, 'ATK': 2, 'DEF': 3, 'AGI': 3, 'EVA': 3, 'WPN': 'stick', 'WM': 2}

you will have a chance to spawn a weaker or stronger version like

{'ID': 1.1, 'HP': 6, 'ATK': 4, 'DEF': 4, 'AGI': 3, 'EVA': 3, 'WPN': 'club', 'WM': 3} -Strong
{'ID': 1.2, 'HP': 2, 'ATK': 1, 'DEF': 3, 'AGI': 1, 'EVA': 3, 'WPN': 'none', 'WM': 2} -Weaker
{'ID': 'b1', 'HP': 15, 'ATK': 7, 'DEF': 9, 'AGI': 2, 'EVA': 2, 'WPN': 'great Club', 'WM': 7} - Boss version

a big thing to keep in mind is its all based off what monsters you feed it, and you can set up boss types as a standalone item that can be used for other variations.


Overall, MonsterGenPro v0.1 could be an invaluable resource for authors, game developers, and tabletop enthusiasts who need to quickly populate their creative works with a diverse array of characters and adversaries. The spreadsheet integration and variation system sound like they would streamline the process and inject an element of unpredictability. I'd be curious to see how this generative tool could enhance the experience for users across different mediums.

I hope to have a full working prototype by Aug 2024
