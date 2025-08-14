const initTechnicalTable=()=>{renderTechnicalTable(),renderTechnicalPagination()},renderTechnicalTable=()=>{const c=document.getElementById("technicalTableBody"),e=(technicalCurrentPage-1)*technicalRowsPerPage,a=Math.min(e+technicalRowsPerPage,technicalTotalRows);c.innerHTML="";for(let n=e;n<a;n++){const t=allTechnicalDataReverse[n],l=document.createElement("tr");l.innerHTML=`
          <td class="price-neutral">${t.full_date}</td>
          <td class="price-neutral">${t.MFI}</td>
          <td class="price-neutral">${t.RSI}</td>
          <td class="price-neutral">${t.MACD}</td>
      `,c.appendChild(l)}document.getElementById("technicalStartEntry").textContent=e+1,document.getElementById("technicalEndEntry").textContent=a,document.getElementById("technicalTotalEntries").textContent=technicalTotalRows};
