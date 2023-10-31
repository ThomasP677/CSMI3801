


def String_compression(word: str) -> str:
    newString = ""
    x = 1
    if len(word) == 0:
        return ""
    if len(word) == 1:
        return string + "1"
    for i in range(len(word)):
        if word[i] == word[i-1]:
            x += 1
        else:
            newString += word[i-1]+ str(x)
            x = 1
    newString += word[i-1]+ str(x)
    return newString
            


if __name__ == "__main__":
    print(String_compression("aabcccccaaa"))
