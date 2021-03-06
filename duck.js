Duck = {
    "Quack": function(){return "quack";}
}

Cat = {
    "Quack": function(){return "meow";}
}

Dog = {
    "Bark": function(){return "bark";}
}

function MakeQuack(x){console.log(x.Quack);}

MakeQuack(Duck);
MakeQuack(Cat);
MakeQuack(Dog); // Prints "undefined" because JS doesn't care about type safety
