window.onload = function(){
  // var bg1 = document.querySelector('#bg1'),
  //     bg1Height = bg1.clientHeight,
  //     bg2 = document.querySelector('#bg2'),
  //     bg2Height = bg2.clientHeight,
  //     coverView = document.querySelector('.cover_area'),
  //     coverHeight = coverView.clientHeight,
  //     topValue = bg1Height + bg2Height,
  //     totalScroll = topValue + coverHeight,
  //     scrollCount = 0;
  //     scrollPosition = 0;
  //     scrollStatus = 0;
  //     stopper = 0;
  // var scrollCounter = function(num){
  //     return scrollCount - num
  // }
  //
  // var topChanger = function(topValue){
  //   coverView.style.top = topValue + 'px';
  // }
  //
  // var getScrollState = function(scrollLength){
  //   if(scrollPosition < scrollLength){
  //     scrollStatus = 0
  //     console.log(scrollStatus);
  //   }else{
  //     scrollStatus = -1;
  //     console.log(scrollStatus);
  //   }
  // }
  //
  // window.addEventListener('scroll', function(e){
  //   var scrollLength = window.scrollY,
  //       currentTop = coverView.style.top,
  //       currentTopValue = currentTop.substring(0, currentTop.length-2) * 1;
  //
  //   if(scrollLength >= bg1Height){
  //     stopper = 1;
  //   }
  //   if(stopper === 1 && scrollStatus === 0){
  //     topChanger(topValue = topValue - 40);
  //     totalScroll =-40;
  //     if(topValue >= bg1Height){
  //     }
  //     return
  //   }
  //   //else if(stopper === 1 && scrollStatus === -1){
  //   //   window.scrollTo(0, bg1Height);
  //   //   topChanger(topValue = topValue + 40);
  //   //   return
  //   // }
  //   getScrollState(scrollLength);
  //   console.log(scrollLength);
  //
  //   scrollPosition = scrollLength;
  //   return
  // });
  var bodyElem = document.body,
      fixedArea = document.querySelector('.fixed_area'),
      coverArea = document.querySelector('.cover_area'),
      bg1 = document.querySelector('#bg1'),

      fixedHeight = fixedArea.clientHeight,
      bg1Height = bg1.clientHeight,
      coverHeight = coverArea.clientHeight,

      totalScroll = fixedHeight + coverHeight,
      beforeScrollPos = 0,
      scrollStatus = 0;

      coverArea.style.top = fixedHeight + 'px';

      var getScrollDirection = function(scrollPos){
        if(beforeScrollPos < scrollPos){
          scrollStatus = 0
        }else{
          scrollStatus = -1;
        }
      }

      bodyElem.style.height = totalScroll + 'px';

      window.addEventListener('scroll', function(e){
        scrollPos = window.pageYOffset;
        getScrollDirection(scrollPos);

        if(scrollPos > 0 && scrollPos < bg1Height){
          fixedArea.style.top = -scrollPos + 'px';
        }else if(scrollPos > bg1Height && scrollStatus === 0){
          fixedArea.style.top = -bg1Height + 'px';
        }else if(scrollPos < 30 && scrollStatus === -1){
          fixedArea.style.top = 0 + 'px';
        }
        beforeScrollPos = scrollPos;
      });

};
