function permPalidrom(phrase){
    //TODO: could use const instead of var, because value doesn't change
    var count = buildCharFrequencyTable(phrase);
    return checkMaxOneOdd(countMap)
}

function checkMaxOneOdd(countMap){
    var count = 0;
    for(i of countMap.values()){
        var count = 0;
        for(i of countMap.values()){
            if(i % 2){
                count += 1
            }
            if(count > 1){
                return false;
            }
        }
    }
    return true;
}

function buildCharFrequencyTable(phrase){
    var mappy = new Map()

    for( c in phrase){
        if(mappy.get(phrase[c]) >= 0){
            mappy.set(phrase[c], mappy.get(pharse[c]+1))
        } else{
            mappy.set(phrase[c],1)
        }
    }
    return mappy;
}
