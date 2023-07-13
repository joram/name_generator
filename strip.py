#!/usr/bin/env python3

male_data = {}
female_data = {}
with open("./17100147.csv", "r") as f:
    lines = f.readlines()
    for line in lines:
        parts = line.split(",")
        sex = parts[3].strip('"')
        name = parts[4].strip('"')
        indicator = parts[5].strip('"')
        value = parts[12].strip('"')
        if indicator == 'Proportion (%)':
            if sex == 'Male':
                if name not in male_data:
                    male_data[name] = 0
                male_data[name] += float(value)
            elif sex == 'Female':
                if name not in female_data:
                    female_data[name] = 0
                female_data[name] += float(value)

with open("./male_data.csv", "w") as f:
    total = 0
    for name, weight in male_data.items():
        f.write(name + "," + str(weight) + "\n")
        total += weight
    print("MALE:", total)

with open("./female_data.csv", "w") as f:
    total = 0
    for name, weight in female_data.items():
        f.write(name + "," + str(weight) + "\n")
        total += weight
    print("FEMALE:", total)