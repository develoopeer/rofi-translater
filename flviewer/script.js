let current_index = 0;
let main_word = document.getElementById("main-word")
let curr_word = document.getElementById('current-value')
let trans_word = document.getElementById("trans-word")
let date = document.getElementById("date")
let max_value = document.getElementById("max-value")

let somedata

window.onload = async function(){
    fetch('dump.json').then( response => response.json()).then(function(result){
        somedata = result  
    }).then(function(){
        change_word_to_index(somedata , 0);
        date.innerHTML = new Date().toLocaleDateString("en-GB").replaceAll("/" , '.')
        curr_word.innerHTML = current_index + 1;
        max_value.innerHTML = Object.keys(somedata).length;
    })

}

function change_word_to_index(words , index){
    main_word.innerHTML = Object.keys(words)[index]
    trans_word.innerHTML = Object.values(words)[index]
}


window.onkeydown = function(key){
    if (key.code == 'ArrowRight'){
        let cv = parseInt(curr_word.innerHTML);
        let mv = parseInt(max_value.innerHTML);
        if (cv + 1 <= mv){
            current_index+= 1
            curr_word.innerHTML = current_index + 1;
            change_word_to_index(somedata , current_index)
        } 
    }
    if (key.code == 'ArrowLeft'){
        let cv = parseInt(curr_word.innerHTML);
        if (cv -1 > 0){
            curr_word.innerHTML = current_index;
            current_index-= 1
            change_word_to_index(somedata , current_index)
        } 
    }
}