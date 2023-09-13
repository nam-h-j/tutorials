window.onload = function(){
    var container = document.querySelector('.matter_canvas');
    
    //1. Loading engine and renderer
    var engine = Matter.Engine.create();
    var render = Matter.Render.create({
        element : container,
        engine : engine,
        options :{
            width: 437,
            heigth: 382,
            wireframes: false,
            background: 'transparent',
            wireframeBackground: 'transparent'
        }
    });

    function __random(strVal, endVal){
        Matter.Common.random(strVal, endVal)
    }
    //2. Adding IMG sources and set object options
    function choco_opt(num){
        return {render:{sprite:{
            texture: '../img/choco' + num + '.png',
            xScale: .9, yScale: .9
        }}}
    }
    function slime_opt(){
        return {render:{sprite:{
            texture: '../img/slime1.png',
            xScale: .9, yScale: .9
        }}}
    }
    var ground_opt = {
        isStatic: true,
        render: { fillStyle: "transparent" }
    }
    
    //3. Make ground
    //3_1 make ground shape as rectangle.
    var grounds = [];
    
    var ground1 = Matter.Bodies.rectangle(290,320,900,1,ground_opt)
    Matter.Body.rotate( ground1, 2.6);
    var ground2 = Matter.Bodies.rectangle(210,355,900,1,ground_opt)
    Matter.Body.rotate( ground2, 0.5);
    var ground3 = Matter.Bodies.rectangle(110,300,900,1,ground_opt)
    Matter.Body.rotate( ground3, 4.1);
    var ground4 = Matter.Bodies.rectangle(320,300,900,1,ground_opt)
    Matter.Body.rotate( ground4, 5.5);
    var ground5 = Matter.Bodies.rectangle(405,180,800,1,ground_opt)
    Matter.Body.rotate( ground5, 1.9);
    var ground6 = Matter.Bodies.rectangle(415,100,800,1,ground_opt)
    Matter.Body.rotate( ground6, 1.5);
    var ground7 = Matter.Bodies.rectangle(30,180,900,1,ground_opt)
    Matter.Body.rotate( ground7, 4.3);
    var ground8 = Matter.Bodies.rectangle(20,100,900,1,ground_opt)
    Matter.Body.rotate( ground8, 4.7);
    var ground9 = Matter.Bodies.rectangle(10,30,900,1,ground_opt)
    Matter.Body.rotate( ground9, 6.25);
    grounds.push(ground1, ground2, ground3, ground4, ground5, ground6, ground7, ground8, ground9);
    //3_1. Draw in the world(canvas)
    Matter.World.add(engine.world,grounds);

    //4. Make my Heart objects 
    //4_1. Get random heart Image mapping for objects
    function getRandomInt(max) {
        return Math.floor(Math.random() * Math.floor(max)) + 1;
    }
    //4_2. Make heart object
    function addChoco(num){
        var choco = Matter.Bodies.circle(Matter.Common.random(120, 320),100,20,choco_opt(num));
        choco.restitution = 0.7;
        choco.frictionAir = 0.001;
        return choco;
    };
    //4_3. Make slime object
    function addSlime(num){
        var choco = Matter.Bodies.circle(Matter.Common.random(100, 250),100,20,slime_opt());
        choco.restitution = 1;
        choco.frictionAir = 0.001;
        return choco;
    };
    //4_4. Draw in the world(canvas)one by object every 0.1s
    var intervals = setInterval(function(){   
        Matter.World.add(engine.world, addChoco(getRandomInt(4)));
    }, 100);
    var intervals_slime = setInterval(function(){   
        Matter.World.add(engine.world, addSlime());
    }, 2500);

    // function setIntervalTime(time){ return time; }
    var animationTime;
    if(status == 30){
        animationTime = 1000;
    }else if(status == 50){
        animationTime = 2600;
    }else if(status == 80){
        animationTime = 4000;
    }else if(status == 100){
        animationTime = 6000;
    }else{
        animationTime = 0;
    }
    setTimeout(function(){
        clearInterval(intervals)
        clearInterval(intervals_slime)
    }, animationTime);   
    
    setTimeout(function(){
        Matter.Render.stop(render);
    }, animationTime + 5000);

    //3. Adding some mouse controlers
    var mouse = Matter.Mouse.create(render.canvas);
    var mouseConstraint = Matter.MouseConstraint.create(engine, {
        mouse: mouse,
        constraint: {
            render: {visible: false}
        }
    });
    render.mouse = mouse;
    Matter.World.add(engine.world,[mouseConstraint]);

    //4. Add odjects in the World and run it
    Matter.Engine.run(engine);
    Matter.Render.run(render);
}