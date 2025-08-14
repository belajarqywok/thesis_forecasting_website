// init table
const initTable = () => {
  renderTable()
  renderPagination()
}


const formatRupiah = (num) => {
  if (num >= 1000) {
    return "Rp " + new Intl.NumberFormat("id-ID", { 
      minimumFractionDigits: 2, 
      maximumFractionDigits: 2 
    }).format(num);
  }
  return num.toString(); // angka kecil tanpa format
}

// render table rows
const renderTable = () => {
  const tableBody  = document.getElementById('tableBody')
  const startIndex = (currentPage - 1) * rowsPerPage
  const endIndex   = Math.min(startIndex + rowsPerPage, totalRows)
  
  tableBody.innerHTML = ''
  
  for (let i = startIndex; i < endIndex; i++) {
    const row = allOHLCData1[i]
    const tr = document.createElement('tr')
      
    tr.innerHTML = `
      <td class="price-neutral">${row.full_date}</td>
      <td class="price-neutral">${formatRupiah(row.open)}</td>
      <td class="price-neutral">${formatRupiah(row.high)}</td>
      <td class="price-neutral">${formatRupiah(row.low)}</td>
      <td class="price-neutral">${formatRupiah(row.close)}</td>
      <td class="price-neutral">${row.volume}</td>
    `
      
    tableBody.appendChild(tr)
  }
  
  // update pagination info
  document.getElementById('startEntry').textContent   = startIndex + 1
  document.getElementById('endEntry').textContent     = endIndex
  document.getElementById('totalEntries').textContent = totalRows
}
