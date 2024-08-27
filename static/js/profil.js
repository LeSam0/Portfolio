function navbar(name) {
    let elements = ['About-Me', 'My-Parcours', 'My-Experience']
    for (element of elements) {
        if (element != name) {
            document.querySelector("."+element).hidden = true;
            document.querySelector("."+element).setAttribute('class',element);
        }
        else {
            document.querySelector("."+element).hidden = false;
            document.querySelector("."+element).setAttribute('class',element+' active');
        }
    }
}

var json = []

async function LoadPage() {
    const response = await fetch('/a_propos/data', { method: 'GET', })
    json = await response.json();
    LoadCard()
} 

function LoadCard(jsonfiltre) {
    jsonfiltre = jsonfiltre == undefined ? json : jsonfiltre
    var container = document.querySelector(".portfolio-content");
    var card = document.querySelector(".card");
    card.hidden = jsonfiltre.length == 0 ? true : false;
    for (var index = 0; index < jsonfiltre.length; index++) {
        card.querySelector('#Image').src = "/static/image/"+jsonfiltre[index].Image;
        card.querySelector('#Name').innerHTML = jsonfiltre[index].Name;
        card.querySelector('#Type').textContent = jsonfiltre[index].Type;
        card.querySelector('#Id').textContent = jsonfiltre[index].Id;
        if (index > 0) {
            container.appendChild(card.cloneNode(true));
        }
    }  
}

function filtre(filtre) {
    filtre = filtre == "" ? "Tous" : filtre;
    let title = ["Tous", "Web", "Application", "GameProgramming", "Cybersécurité"];
    for (let index = 0; index < title.length; index++) {
        document.querySelector('#'+title[index]).innerHTML = filtre == title[index] ? "<u>"+title[index]+"</u>" : title[index];
    }
    let jsonfiltre = filtre == "Tous" ? undefined : json.filter((data) => data.Type == filtre);
    LoadCard(jsonfiltre)
}