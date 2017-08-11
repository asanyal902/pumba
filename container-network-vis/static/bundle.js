(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
function populate() {
    fetch('/containers/json').then(function(r) {
        return r.json()
    }).then(function(j) {
        console.log("json: ", j);

        var table = document.getElementById("containerTable");
        for (var i = 0; i < j.length; i++) {
            var name = j[i].Names[0].substring(1);
            console.log("name: ", j[i].Names[0]);
            var row = table.insertRow();
            var cell1 = row.insertCell(0);
            var cell2 = row.insertCell(1);
            var cell3 = row.insertCell(2);
            var cell4 = row.insertCell(3);
            var cell5 = row.insertCell(4);
            cell1.innerHTML = name;
            cell2.innerHTML = "-";
            cell3.innerHTML = "-";
            cell4.innerHTML = "-";
            cell5.innerHTML = "-";
        }
    });
}

function enter(){
    // var xhttp = new XMLHttpRequest();
    // xhttp.open("POST", "localhost:8009");
    // xhttp.setRequestHeader("Content-type", "application/json");
    console.log("entered");
     var pumbaObject = new Object();
     var containerName = document.getElementById("containerName").value;
     if (containerName !== null && containerName !== "") {
         pumbaObject["containerName"] = containerName;
     }
     var delay = document.getElementById("delay").value;
     if (delay !== null && delay !== "") {
         pumbaObject["delay"] = parseInt(delay);
     }
     var loss = document.getElementById("loss").value;
     if (loss !== null && loss !== "") {
         pumbaObject["loss"] = parseInt(loss);
     }
     var rate = document.getElementById("rate").value;
     if (rate !== null && rate !== "") {
         pumbaObject["rate"] = parseInt(rate);
     }
     var time = document.getElementById("time").value;
     if (time !== null && time !== "") {
         pumbaObject["time"] = time + "m";
     }
    // var pumbaObject = `{containerName: ` + containerName
    //     + `, delay:` + delay + `, loss:` + loss
    //     +  `, rate: ` + rate + `, time:`+ time +`m }`;
    // xhttp.send(pumbaObject);

    var table = document.getElementById("containerTable");
    console.log("table: ", table);
    for (var i = 1; i < table.rows.length; i++) {
        var cells = table.rows[i].cells;
        console.log("cells: ", cells);
        console.log("cells length: ", cells.length);
        console.log("cells[0]: ", cells[0].innerHTML);
        console.log("containerName: ", containerName);
        if (cells[0].innerHTML === containerName && cells[4].innerHTML === "-") {
            cells[1].innerHTML = delay;
            cells[2].innerHTML = loss;
            cells[3].innerHTML = rate;
            var d = new Date(time * 1000 * 60);
            initializeClock(cells, d);

            var xhttp = new XMLHttpRequest();
            xhttp.open("POST", "http://localhost:8009/netem");
            xhttp.setRequestHeader("Content-type", "application/json");


            xhttp.send(JSON.stringify(pumbaObject));
        }
    }
}

function initializeClock(cells, endtime) {
    var s = new Date(1000);

    function updateClock(endtime) {
        endtime = endtime - s;
        var seconds = Math.floor((endtime / 1000) % 60);
        var minutes = Math.floor((endtime / 1000 / 60) % 60);

        var minutesSpan = ('0' + minutes).slice(-2);
        var secondsSpan = ('0' + seconds).slice(-2);

        cells[4].innerHTML = minutesSpan + ":" + secondsSpan;

        if (endtime > 0) {
            setTimeout(updateClock, 1000, endtime);
        } else {
            cells[4].innerHTML = "-";
            cells[1].innerHTML = "-";
            cells[2].innerHTML = "-";
            cells[3].innerHTML = "-";
        }
    }
    updateClock(endtime);
}


populate();

var enterBtn = document.getElementById("enter");

enterBtn.onclick = function(ev){
    enter();
};

},{}]},{},[1]);
