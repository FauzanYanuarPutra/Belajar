class Coding {
  #model = "";
  #make = "";

  constructor(make, model) {
    this.#make = make;
    this.#model = model;
  }

  get show() {
    return `Saya belajar bahasa ${this.#make}`;
  }

  get withFramework() {
    return `${this.show}, dengan framework ${this.#model}`;
  }

  set withFramework(value) {
    if (value === "") {
      console.log("the make cannot be an empty string!");
      return;
    }
    this.#make = value;
  }
}

// bahasa, framework

let data1 = new Coding("PHP", "Laravel");
console.log(data1.show);
console.log(data1.withFramework);
