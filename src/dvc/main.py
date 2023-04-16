import os
import matplotlib.pyplot as plt
import cv2

test_image = cv2.imread('images/logo.jpg')
plt.imshow(test_image)
plt.title('Image created out of patches \n')
plt.show()
