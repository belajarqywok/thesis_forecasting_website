// pagination
function renderTechnicalPagination() {
  const totalPages = Math.ceil(technicalTotalRows / technicalRowsPerPage);
  const pageNumbers = document.getElementById('technicalPageNumbers');
  const prevBtn = document.getElementById('technicalPrevBtn');
  const nextBtn = document.getElementById('technicalNextBtn');
  
  pageNumbers.innerHTML = '';
  
  let startPage = Math.max(1, technicalCurrentPage - 2);
  let endPage = Math.min(totalPages, technicalCurrentPage + 2);
  
  if (technicalCurrentPage <= 3) {
      endPage = Math.min(5, totalPages);
  }
  if (technicalCurrentPage >= totalPages - 2) {
      startPage = Math.max(totalPages - 4, 1);
  }
  
  if (startPage > 1) {
      addTechnicalPageButton(1);
      if (startPage > 2) {
          addTechnicalEllipsis();
      }
  }
  
  for (let i = startPage; i <= endPage; i++) {
      addTechnicalPageButton(i);
  }
  
  if (endPage < totalPages) {
      if (endPage < totalPages - 1) {
          addTechnicalEllipsis();
      }
      addTechnicalPageButton(totalPages);
  }
  
  prevBtn.disabled = technicalCurrentPage === 1;
  nextBtn.disabled = technicalCurrentPage === totalPages;
}

function addTechnicalPageButton(pageNum) {
  const pageNumbers = document.getElementById('technicalPageNumbers');
  const button = document.createElement('button');
  button.className = `pagination-btn ${pageNum === technicalCurrentPage ? 'active' : ''}`;
  button.textContent = pageNum;
  button.onclick = () => goToTechnicalPage(pageNum);
  pageNumbers.appendChild(button);
}

function addTechnicalEllipsis() {
  const pageNumbers = document.getElementById('technicalPageNumbers');
  const ellipsis = document.createElement('span');
  ellipsis.className = 'pagination-btn';
  ellipsis.textContent = '...';
  ellipsis.style.cursor = 'default';
  ellipsis.style.background = 'transparent';
  ellipsis.style.border = 'none';
  pageNumbers.appendChild(ellipsis);
}

function goToTechnicalPage(pageNum) {
  technicalCurrentPage = pageNum;
  renderTechnicalTable();
  renderTechnicalPagination();
}

function changeTechnicalPage(direction) {
  const totalPages = Math.ceil(technicalTotalRows / technicalRowsPerPage);
  const newPage = technicalCurrentPage + direction;
  
  if (newPage >= 1 && newPage <= totalPages) {
      goToTechnicalPage(newPage);
  }
}

function changeTechnicalRowsPerPage() {
  const select = document.getElementById('technicalRowsPerPage');
  technicalRowsPerPage = parseInt(select.value);
  technicalCurrentPage = 1;
  renderTechnicalTable();
  renderTechnicalPagination();
}