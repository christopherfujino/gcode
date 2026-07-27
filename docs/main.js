const re = /total filament used.*= (.*)$/;

var _makeUpdate = (output) => function(filename, amount) {
  const tr = document.createElement("tr");
  const td0 = document.createElement("td");
  td0.textContent = filename;
  tr.appendChild(td0);
  const td1 = document.createElement("td");
  td1.textContent = `${amount}g`;
  tr.appendChild(td1);
  output.appendChild(tr);
}

var update = null;

function handleFiles() {
  const fileList = this.files;
  for (const file of fileList) {
    file.text().then(function(text) {
      text.split("\n").forEach(function(line) {
        if (line == "") {
          return;
        }
        switch (line[0]) {
          case "G":
            return;
          case "M":
            return;
          case ";":
            // A comment
            const match = line.match(re);
            if (match != null) {
              update(file.name, match[1]);
            }
            return;
          default:
            throw `Foolies! "${line}"`;
        }
      });
    });
  }
}

document.addEventListener("DOMContentLoaded", () => {
  update = _makeUpdate(document.getElementById("output"));

  const inputElement = document.getElementById("input");
  inputElement.addEventListener("change", handleFiles);
});
