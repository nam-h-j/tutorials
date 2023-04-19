# pyautogui
- 마우스 및 키보드를 자동으로 제어 할 수 있는 라이브러리

## Install
```bash
pip install pyautogui
```

## init
```python
import pyautogui

pag = pyautogui

print(pag.size()) # will shows your device monitor resolution
```


## mouse functions 마우스 제어
### 커서위치이동
```pyton
>>> pyautogui.moveTo(x, y, duration=num_seconds)  # move mouse to XY coordinates over num_second seconds
>>> pyautogui.moveRel(xOffset, yOffset, duration=num_seconds)  # move mouse relative to its current position
```
### click
```pyton
pag.click()
```
