#!/usr/bin/env python3

male_data = {}
female_data = {}
with open("./firstnames.csv", "r") as f:
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

surnames = {}
# name,rank,count,prop100k,cum_prop100k,pctwhite,pctblack,pctapi,pctaian,pct2prace,pcthispanic
with open("./surnames.csv", "r") as f:
    lines = f.readlines()
    for line in lines[1:]:
        parts = line.split(",")
        name = parts[0].strip('"')
        count = parts[2].strip('"')
        if name not in surnames:
            surnames[name] = 0
        surnames[name] += float(count)



with open("./name_generator/data/male_names.go", "w") as f:
    f.write("package data\n\n")
    f.write("var MaleNames = map[string]int{\n")
    min_weight = min(male_data.values())
    for name, weight in male_data.items():
        weight = max(1, int(weight/min_weight))
        name = name.capitalize()
        f.write('\t"' + name + '": ' + str(weight) + ",\n")
    f.write("}\n")

with open("./name_generator/data/female_names.go", "w") as f:
    f.write("package data\n\n")
    f.write("var FemaleNames = map[string]int{\n")
    min_weight = min(female_data.values())
    for name, weight in female_data.items():
        weight = max(1, int(weight/min_weight))
        name = name.capitalize()
        f.write('\t"' + name + '": ' + str(weight) + ",\n")
    f.write("}\n")

with open("./name_generator/data/surnames.go", "w") as f:
    f.write("package data\n\n")
    f.write("var Surnames = map[string]int{\n")
    min_weight = min(surnames.values())
    for name, weight in surnames.items():
        weight = max(1, int(weight/min_weight))
        name = name.capitalize()
        f.write('\t"' + name + '": ' + str(weight) + ",\n")
    f.write("}\n")
