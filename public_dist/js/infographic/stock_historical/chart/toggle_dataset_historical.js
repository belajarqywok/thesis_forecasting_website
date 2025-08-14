const toggleDataset=t=>{const e=document.getElementById(t+"-check"),c=e.parentElement;e.checked=!e.checked,e.checked?c.classList.add("checked"):c.classList.remove("checked"),updateChart()};
