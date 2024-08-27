    var json = []

    async function LoadPage() {
        let index = 1;
        const response = await fetch('/mes_experience_professionel/data', { method: 'GET', })
        json = await response.json();
        let indexJson = [0, 1]
        LoadContent(index);
        LoadCaroussel(indexJson);
    } 

    function LoadCaroussel(indexJson) {
        const cards = document.querySelectorAll('.card');
        cards.forEach((card, index) => {
            if (json[2] == undefined && index == 2) {
                card.querySelector('#CarrouselImage'+[index]).src = "/static/image/"+json[indexJson[0]].Image;
                card.querySelector('#CarrouselName'+[index]).innerHTML = json[indexJson[0]].Name;
                card.querySelector('#CarrouselContract'+[index]).textContent = json[indexJson[0]].Contract;
                card.querySelector('#Id'+[index]).textContent = json[indexJson[0]].Id;
            }
            card.querySelector('#CarrouselImage'+[index]).src = "/static/image/"+json[indexJson[index]].Image;
            card.querySelector('#CarrouselName'+[index]).innerHTML = json[indexJson[index]].Name;
            card.querySelector('#CarrouselContract'+[index]).textContent = json[indexJson[index]].Contract;
            card.querySelector('#Id'+[index]).textContent = json[indexJson[index]].Id;
        });
    }

    function LoadContent(index) {
        document.querySelector('#Description').innerHTML = json[index].Description;
        document.querySelector('#Experiences').innerHTML = json[index].Experience;
        document.querySelector('#Creation').innerHTML = json[index].Creation;
        document.querySelector('#Createur').innerHTML = json[index].Creator;
        document.querySelector('#Dirigeant').innerHTML = json[index].Dirigeant;
        document.querySelector('#Localisation').innerHTML = json[index].Localisation;
        document.querySelector('#Activite').innerHTML = json[index].Activite;
        document.querySelector('#Name').innerHTML = json[index].Name;
        document.querySelector('#Contract').innerHTML = json[index].Contract;
        document.querySelector('#Lien').href = json[index].Lien;
        document.querySelector('#Image').src = "/static/image/"+json[index].Image;
    }

    function switchgauche() {
        let indexbefore = (document.querySelector('#Id0').textContent)-1
        let indexselect = (document.querySelector('#Id1').textContent)-1
        LoadContent(indexbefore)
        let indexJson = [indexselect, indexbefore]
        LoadCaroussel(indexJson)
    }

    function switchdroit() {
        let indexnext = (document.querySelector('#Id2').textContent)-1
        let indexselect = (document.querySelector('#Id1').textContent)-1
        LoadContent(indexnext)
        let indexJson = [indexselect, indexnext]
        LoadCaroussel(indexJson)
    }