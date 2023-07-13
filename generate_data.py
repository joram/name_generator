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

with open("./name_generator/data.go", "w") as f:
    f.write("package name_generator\n\n")
    f.write("var MaleNames = map[string]int{\n")
    min_weight = min(male_data.values())
    for name, weight in male_data.items():
        weight = max(1, int(weight/min_weight))
        f.write('\t"' + name + '": ' + str(weight) + ",\n")
    f.write("}\n\n")

    f.write("var FemaleNames = map[string]int{\n")
    min_weight = min(female_data.values())
    for name, weight in female_data.items():
        weight = max(1, int(weight/min_weight))
        f.write('\t"' + name + '": ' + str(weight) + ",\n")
    f.write("}\n")
