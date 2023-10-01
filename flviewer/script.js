let current_index = 0;
let main_word = document.getElementById("main-word")
let curr_word = document.getElementById('current-value')
let trans_word = document.getElementById("trans-word")
let date = document.getElementById("date")
let max_value = document.getElementById("max-value")
let audio = document.getElementById('media');

let somedata

window.onload = async function(){
    curr_word.value = 1;
    fetch('another.json').then( response => response.json()).then(function(result){
        somedata = result  
    }).then(function(){
        change_word_to_index(somedata , 0);
        audio.src = 'data:audio/wav;base64,' + Object.values(somedata)[0][0]
        date.innerHTML = new Date().toLocaleDateString("en-GB").replaceAll("/" , '.')
        curr_word.innerHTML = current_index + 1;
        max_value.innerHTML = Object.keys(somedata).length;
    })

}

function change_word_to_index(words , index){
    main_word.innerHTML = Object.keys(words)[index]
    trans_word.innerHTML = Object.values(words)[index][1]
}


window.onkeydown = function(key){
    if (key.code == 'ArrowRight'){
        let cv = parseInt(curr_word.value);
        let mv = parseInt(max_value.innerHTML);
        if (cv + 1 <= mv){
            current_index+= 1
            curr_word.value = current_index + 1;
            change_word_to_index(somedata , current_index)
            audio.src = 'data:audio/wav;base64,' + Object.values(somedata)[current_index][0]
        }else{
            current_index = 0
            curr_word.value = current_index + 1;
            change_word_to_index(somedata , current_index)
            audio.src = 'data:audio/wav;base64,' + Object.values(somedata)[current_index][0]
        }
    }
    if (key.code == 'ArrowLeft'){
        let cv = parseInt(curr_word.value);
        let mv = parseInt(max_value.innerHTML);
        if (cv -1 > 0){
            curr_word.value = current_index;
            current_index-= 1
            change_word_to_index(somedata , current_index)
            audio.src = 'data:audio/wav;base64,' + Object.values(somedata)[current_index][0]
        }else{
            current_index = mv - 1
            curr_word.value = current_index + 1;
            change_word_to_index(somedata , current_index)
            audio.src = 'data:audio/wav;base64,' + Object.values(somedata)[current_index+1][0]
        }
    }
}

document.getElementById('current-value').onchange = function(){
    current_index = document.getElementById('current-value').value
    if (current_index > parseInt(max_value.innerHTML)){
        current_index = parseInt(max_value.innerHTML)
        curr_word.value = current_index;
    }
    change_word_to_index(somedata , current_index - 1)
    audio.src = 'data:audio/wav;base64,' + Object.values(somedata)[current_index-1][0]
}

document.getElementById('play').onclick = function(){
    // audio.play()
    // 'data:audio/wav;base64,'
    if (audio.paused){
        audio.play();
    }else{
        audio.pause();
    }
}