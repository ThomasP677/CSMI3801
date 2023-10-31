
const palit = "tacocat"

const palidrome = {
    pal_reverse: (pal.split("").reverse().join("")).toString(),
    Palidrome: function(pal) {
      return pal == this.pal_reverse;
    }
  };

name = palidrome.Palidrome(palit);
console.log(name);

