import json
import random
import urllib.request

url = urllib.request.urlopen("https://raw.githubusercontent.com/sindresorhus/mnemonic-words/master/words.json")
words = json.loads(url.read())
random_word = random.choice(words)
for x in range(50):
    filename = "file_"+str(x)+".txt"
    fo = open(filename, "w")
    out = ""
    for y in range(100000):
        if y%10==0:
            out=out+"\n"
        out = out+random.choice(words)+" "

    fo.write(out)

    # Close opend file
    fo.close()
