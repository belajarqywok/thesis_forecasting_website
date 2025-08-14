// init technical table
const initTechnicalTable = () => {
  renderTechnicalTable()
  renderTechnicalPagination()
}

// render technical table
const renderTechnicalTable = () => {
  const tableBody = document.getElementById('technicalTableBody');
  const startIndex = (technicalCurrentPage - 1) * technicalRowsPerPage;
  const endIndex = Math.min(startIndex + technicalRowsPerPage, technicalTotalRows);
  
  tableBody.innerHTML = '';
  
  for (let i = startIndex; i < endIndex; i++) {
      const row = allTechnicalDataReverse[i];
      const tr = document.createElement('tr');
      
      tr.innerHTML = `
          <td class="price-neutral">${row.full_date}</td>
          <td class="price-neutral">${row.MFI}</td>
          <td class="price-neutral">${row.RSI}</td>
          <td class="price-neutral">${row.MACD}</td>
      `;
      
      tableBody.appendChild(tr);
  }
  
  document.getElementById('technicalStartEntry').textContent = startIndex + 1;
  document.getElementById('technicalEndEntry').textContent = endIndex;
  document.getElementById('technicalTotalEntries').textContent = technicalTotalRows;
}