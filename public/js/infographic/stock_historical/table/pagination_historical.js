// render pagination controls
const renderPagination = () => {
  const totalPages = Math.ceil(totalRows / rowsPerPage)
  const pageNumbers = document.getElementById('pageNumbers')

  const prevBtn = document.getElementById('prevBtn')
  const nextBtn = document.getElementById('nextBtn')
  
  // clear existing page numbers
  pageNumbers.innerHTML = ''
  
  // calculate page range to show
  let startPage = Math.max(1, currentPage - 2)
  let endPage = Math.min(totalPages, currentPage + 2)
  
  // adjust range if we're near the beginning or end
  if (currentPage <= 3) {
    endPage = Math.min(5, totalPages)
  }
  if (currentPage >= totalPages - 2) {
    startPage = Math.max(totalPages - 4, 1)
  }
  
  // add first page and ellipsis if needed
  if (startPage > 1) {
    addPageButton(1)
    if (startPage > 2) {
      addEllipsis()
    }
  }
  
  // add page numbers
  for (let i = startPage; i <= endPage; i++) {
    addPageButton(i)
  }
  
  // add ellipsis and last page if needed
  if (endPage < totalPages) {
    if (endPage < totalPages - 1) {
      addEllipsis()
    }
    addPageButton(totalPages)
  }
  
  // update navigation buttons
  prevBtn.disabled = currentPage === 1
  nextBtn.disabled = currentPage === totalPages
}


// add page button
const addPageButton = (pageNum) => {
  const pageNumbers = document.getElementById('pageNumbers')
  const button = document.createElement('button')

  button.className = `pagination-btn ${pageNum === currentPage ? 'active' : ''}`
  button.textContent = pageNum
  button.onclick = () => goToPage(pageNum)

  pageNumbers.appendChild(button)
}


// add ellipsis
const addEllipsis = () => {
  const pageNumbers = document.getElementById('pageNumbers')
  const ellipsis = document.createElement('span')

  ellipsis.className        = 'pagination-btn'
  ellipsis.textContent      = '...'
  ellipsis.style.cursor     = 'default'
  ellipsis.style.background = 'transparent'
  ellipsis.style.border     = 'none'

  pageNumbers.appendChild(ellipsis)
}


// go to specific page
const goToPage = (pageNum) => {
  currentPage = pageNum

  renderTable()
  renderPagination()
}


// change page (previous / next)
const changePage = (direction) => {
  const totalPages = Math.ceil(totalRows / rowsPerPage)
  const newPage    = currentPage + direction
  
  if (newPage >= 1 && newPage <= totalPages) {
    goToPage(newPage)
  }
}


// change rows per page
const changeRowsPerPage = () => {
  const select = document.getElementById('rowsPerPage')
  rowsPerPage = parseInt(select.value)

  // reset to first page
  currentPage = 1

  renderTable()
  renderPagination()
}
