
function creationDateSlide() {
    const print = document.getElementById("rangeValue")
    const change = document.getElementById("CreationDate")
    print.innerText = change.value;

    const form = document.getElementById("filterBox")

    console.log(change.value)

    for (i = 0; i < 52; i++) {
        li = document.getElementById(`box${i + 1}`)
        if (albumDate[i].includes(change.value)) {
            li.style.display = "relative";
        } else {
            li.style.display = "none";
        }
    }
}

function firstAlbumSlide() {
    const print = document.getElementById("firstAlbum")
    const change = document.getElementById("firstAlbumSlider")
    print.innerText = change.value;

    const form = document.getElementById("filterBox")

    for (i = 0; i < 52; i++) {
        li = document.getElementById(`box${i + 1}`)
        if (creationDate[i].includes(change.value)) {
            li.style.display = "relative";
        } else {
            li.style.display = "none";
        }
    }
}

const searchInput = document.querySelector("[data-search]")

let bandName = []
let membersName = []
let albumDate = []
let creationDate = []


objApi.forEach(artist => {
    bandName.push( A.Name) 
    membersName.push( A.Members) 
    albumDate.push(A.FirstAlbum) 
    creationDate.push(A.creationDate.toString())
});

const artistsCards = [];
for (i = 1; i < 53; i++) {
    li = document.getElementById(`box${i}`)
    artistsCards.push(li)
}

searchInput.addEventListener("input", e => {
    const value = e.target.value.toLowerCase()
    for (i = 0; i < 52; i++) {
        li = document.getElementById(`box${i + 1}`)
        if (bandName[i].includes(value)) {
            li.style.display = "flex";
        } else if (albumDate[i].includes(value)) {
            li.style.display = "flex";
        } else if (creationDate[i].includes(value)) {
            li.style.display = "flex";
        } else if ((membersName[i].join(' ').toLowerCase()).includes(value)) {
            li.style.display = "flex";
        } else {
            li.style.display = "none";
        }
    }
})

function validate() {
    displayNothing();
    for (i = 1; i < 8; i++) {
        if (document.getElementById('checkbox' + i).checked) {
            for (j = 0; j < 52; j++) {
                li = document.getElementById(`box${j + 1}`)
                if (membersName[j].length == i) {
                    console.log(membersName[j].length);
                    console.log(membersName[j])
                    li.style.display = "flex";
                }
            }
        }
    }
}

function displayNothing() {
    for (i = 1; i < 53; i++) {
        document.getElementById('box' + i).style.display = "none";
    }
}