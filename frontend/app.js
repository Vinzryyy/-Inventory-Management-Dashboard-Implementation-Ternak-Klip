const inventoryBody = document.getElementById("inventory-body");
const feedback = document.getElementById("feedback");
const refreshButton = document.getElementById("refresh-button");
const productCount = document.getElementById("product-count");
const apiStatus = document.getElementById("api-status");

async function loadProducts() {
  setLoadingState(true);

  try {
    const response = await fetch("/api/v1/products", {
      headers: {
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`HTTP ${response.status}`);
    }

    const result = await response.json();
    const products = Array.isArray(result.data) ? result.data : [];

    renderProducts(products);
    feedback.textContent = `Data berhasil dimuat pada ${new Date().toLocaleString("id-ID")}.`;
    feedback.classList.remove("error");
    apiStatus.textContent = "Online";
    productCount.textContent = String(products.length);
  } catch (error) {
    inventoryBody.innerHTML = `
      <tr>
        <td colspan="3" class="empty-state">Gagal mengambil data dari backend.</td>
      </tr>
    `;
    feedback.textContent = "Tidak dapat terhubung ke endpoint inventory. Pastikan backend dan database aktif.";
    feedback.classList.add("error");
    apiStatus.textContent = "Offline";
    productCount.textContent = "0";
    console.error(error);
  } finally {
    setLoadingState(false);
  }
}

function renderProducts(products) {
  if (products.length === 0) {
    inventoryBody.innerHTML = `
      <tr>
        <td colspan="3" class="empty-state">Belum ada data untuk ditampilkan.</td>
      </tr>
    `;
    return;
  }

  inventoryBody.innerHTML = products
    .map((product) => {
      const updatedAt = formatTimestamp(product.last_updated);
      return `
        <tr>
          <td>${escapeHtml(product.name)}</td>
          <td><span class="stock-pill">${product.stock_count}</span></td>
          <td>${updatedAt}</td>
        </tr>
      `;
    })
    .join("");
}

function formatTimestamp(value) {
  if (!value) {
    return "-";
  }

  const date = new Date(value);
  if (Number.isNaN(date.getTime())) {
    return value;
  }

  return new Intl.DateTimeFormat("id-ID", {
    dateStyle: "medium",
    timeStyle: "short",
  }).format(date);
}

function setLoadingState(isLoading) {
  refreshButton.disabled = isLoading;
  refreshButton.textContent = isLoading ? "Memuat..." : "Refresh Data";
  if (isLoading) {
    apiStatus.textContent = "Memuat";
  }
}

function escapeHtml(value) {
  return String(value)
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#39;");
}

refreshButton.addEventListener("click", loadProducts);
document.addEventListener("DOMContentLoaded", loadProducts);
