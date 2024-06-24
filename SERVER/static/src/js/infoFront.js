const infoBloc = document.getElementById('infoBloc');
const infoButton = document.getElementById('moreInfos');


infoButton.onclick = function () {
  if (infoBloc.style.display === 'block') {
    infoBloc.style.display = 'none';
  } else {
    infoBloc.style.display = 'block';
  };
};
