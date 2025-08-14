const initTable=()=>{renderTable(),renderPagination()},formatRupiah=t=>t>=1e3?"Rp "+new Intl.NumberFormat("id-ID",{minimumFractionDigits:2,maximumFractionDigits:2}).format(t):t.toString(),renderTable=()=>{const t=document.getElementById("tableBody"),n=(currentPage-1)*rowsPerPage,a=Math.min(n+rowsPerPage,totalRows);t.innerHTML="";for(let r=n;r<a;r++){const e=allOHLCData1[r],o=document.createElement("tr");o.innerHTML=`
      <td class="price-neutral">${e.full_date}</td>
      <td class="price-neutral">${formatRupiah(e.open)}</td>
      <td class="price-neutral">${formatRupiah(e.high)}</td>
      <td class="price-neutral">${formatRupiah(e.low)}</td>
      <td class="price-neutral">${formatRupiah(e.close)}</td>
      <td class="price-neutral">${e.volume}</td>
    `,t.appendChild(o)}document.getElementById("startEntry").textContent=n+1,document.getElementById("endEntry").textContent=a,document.getElementById("totalEntries").textContent=totalRows};
