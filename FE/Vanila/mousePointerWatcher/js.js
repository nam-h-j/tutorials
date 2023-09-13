window.onload = function(){
  var eyeball = function(number){
    var eye = document.querySelector('.eye'+number),
        pupil = eye.querySelector('.pupil');
        eyearea = eye.getBoundingClientRect();

    var spin = function(mouseX, mouseY){
      var eye_centerY = mouseY - (eyearea.y + eyearea.height * 0.5),
          eye_centerX = mouseX - (eyearea.x + eyearea.width * 0.5),
          radian = Math.atan2(eye_centerY, eye_centerX),
          degree = (180*radian/Math.PI);
      pupil.style.transform = 'rotate('+ (degree+90) +'deg)';
    }
    return{
      spin : spin
    }
  }

  var eye1 = eyeball(1);
  var eye2 = eyeball(2);

  window.addEventListener('mousemove', function(e){
    eye1.spin(e.pageX, e.pageY);
    eye2.spin(e.pageX, e.pageY);
  })
}
