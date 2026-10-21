const urlInput = document.querySelector('#url');
const status = document.querySelector('#status');
const historyList = document.querySelector('#history');
const clearButton = document.querySelector('#clear');
const storageKey = 'sniplink-session-links';
let toastTimer;

function validURL(value) {
  try {
    return ['http:', 'https:'].includes(new URL(value).protocol);
  } catch {
    return false;
  }
}

function readHistory() {
  try {
    const items = JSON.parse(sessionStorage.getItem(storageKey) || '[]');
    return Array.isArray(items)
      ? items.filter(item => item && validURL(item.shortUrl) && validURL(item.longUrl)).slice(0, 20)
      : [];
  } catch {
    return [];
  }
}

function toast(text) {
  const element = document.querySelector('#toast');
  clearTimeout(toastTimer);
  element.textContent = text;
  element.hidden = false;
  toastTimer = setTimeout(() => { element.hidden = true; }, 2400);
}

async function copyText(text, button) {
  try {
    await navigator.clipboard.writeText(text);
    toast('Link copiado. Pode compartilhar!');
    button.textContent = 'Copiado ✓';
    setTimeout(() => { button.textContent = 'Copiar link'; }, 2200);
  } catch {
    status.textContent = 'Não foi possível acessar a área de transferência. Selecione o link para copiá-lo.';
    status.classList.add('error');
  }
}

function renderHistory() {
  const items = readHistory();
  historyList.replaceChildren();
  document.querySelector('#count').textContent = items.length;
  clearButton.hidden = !items.length;
  if (!items.length) {
    const empty = document.createElement('div');
    empty.className = 'empty';
    empty.innerHTML = '<svg viewBox="0 0 40 40" fill="none" stroke="currentColor" stroke-width="1.6" aria-hidden="true"><rect x="6" y="16" width="20" height="11" rx="5.5" transform="rotate(-35 16 21)"/><rect x="17" y="12" width="20" height="11" rx="5.5" transform="rotate(-35 27 17)"/></svg><span>Seu primeiro atalho começa ali em cima.<br>Vamos encurtar uma boa ideia?</span>';
    historyList.append(empty);
    return;
  }
  items.forEach(item => {
    const row = document.createElement('div');
    row.className = 'link-row';
    const icon = document.createElement('div');
    icon.className = 'link-icon';
    icon.textContent = '↗';
    icon.setAttribute('aria-hidden', 'true');
    const info = document.createElement('div');
    info.className = 'link-info';
    const link = document.createElement('a');
    link.href = item.shortUrl;
    link.target = '_blank';
    link.rel = 'noopener noreferrer';
    link.textContent = item.shortUrl.replace(/^https?:\/\//, '');
    const destination = document.createElement('p');
    destination.textContent = item.longUrl;
    destination.title = item.longUrl;
    const copy = document.createElement('button');
    copy.className = 'copy';
    copy.type = 'button';
    copy.textContent = 'Copiar link';
    copy.addEventListener('click', () => copyText(item.shortUrl, copy));
    info.append(link, destination);
    row.append(icon, info, copy);
    historyList.append(row);
  });
}

const resultURL = document.querySelector('#result-url');
if (resultURL) {
  const item = {
    shortUrl: resultURL.href,
    longUrl: document.querySelector('#result-destination').textContent,
    createdAt: Date.now(),
  };
  try {
    const items = readHistory().filter(entry => entry.shortUrl !== item.shortUrl);
    sessionStorage.setItem(storageKey, JSON.stringify([item, ...items].slice(0, 20)));
  } catch {}
  document.querySelector('#copy-result').addEventListener('click', event => {
    copyText(item.shortUrl, event.currentTarget);
  });
}

document.querySelector('#example').addEventListener('click', () => {
  urlInput.value = 'https://www.wikipedia.org/';
  urlInput.focus();
});
clearButton.addEventListener('click', () => {
  if (!confirm('Limpar o histórico deste navegador? Os links continuam funcionando.')) return;
  try { sessionStorage.removeItem(storageKey); } catch {}
  renderHistory();
  toast('Histórico limpo.');
});
document.querySelector('#shorten-form').addEventListener('submit', () => {
  status.classList.remove('error');
  status.textContent = 'Criando seu atalho…';
  document.querySelector('#submit').disabled = true;
});
renderHistory();

if (!matchMedia('(prefers-reduced-motion: reduce)').matches) {
  let ticking = false;
  addEventListener('scroll', () => {
    if (ticking) return;
    ticking = true;
    requestAnimationFrame(() => {
      document.querySelector('.art').style.transform = `translateY(${Math.min(scrollY * .12, 45)}px)`;
      ticking = false;
    });
  }, { passive: true });
}
