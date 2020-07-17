
function wait(time) {
  return new Promise(resolve => {
    setTimeout(() => {
      console.log("hey")
      resolve('');
    }, time);
  });
}


function inf() {
  wait(2000).then(() => inf())
}

inf()