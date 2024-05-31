import random

def generate_monster(id):
    return {
        'ID': id,
        'HP': random.randint(4, 12),
        'ATK': random.randint(2, 4),
        'DEF': random.randint(2, 4),
        'AGI': random.randint(1, 4),
        'EVA': random.randint(0, 4),
        'WPN': f'Weapon{id}'  # Just as a placeholder name for the weapon
    }

def get_random_monster(monsters):
    return random.choice(monsters)

# Generate a list of 5 monsters
monsters = [generate_monster(i+1) for i in range(5)]

def example_goblin():
    goblin = {
        'ID': 1,
        'HP': 5,
        'ATK': 2,
        'DEF': 3,
        'AGI': 3,
        'EVA': 3,
        'WPN': "stick",
        'WM': 2
    }
    return goblin


# Print all monsters
for monster in monsters:
    print(monster)

# Get a random monster
random_monster = get_random_monster(monsters)
print("\nRandomly selected monster:")
print(random_monster)

sampleGob = example_goblin()

print(sampleGob)