import vk_api,os,random, datetime
from vk_api import VkUpload
from time import sleep


def has_image_extention(filename):
	filename = filename.lower()

	image_extentions = ["jpg","png","bmp","jpeg"]
	for image_extention in image_extentions:
		if image_extention in filename:
			return True

	return False

def sendImages(img_dir = ".", ID = 0, IS_GROUP_CHAT = False, TOKEN = ""):
	vk = vk_api.VkApi(token = TOKEN).get_api()

	files = os.listdir(img_dir)
	files.sort()

	counter = 1
	for filename in files:

		if has_image_extention(filename) == False:
			# file is probably not an image, so skipping
			continue

		# full path to the image
		path_to_image = img_dir + filename

		# upload via API
		upload = VkUpload(vk)
		upload_img = upload.photo_messages(photos = path_to_image)[0]

		# each message will contain "Counter : {number_of_}"
		MESSAGE = "▶ Counter  ┃{}┃ The next appr. at ~ ┃{}┃ ◀".format(counter,
		 datetime.datetime.now() + datetime.timedelta(seconds=4))
		# sending
		print("• {}: Sending {}...".format(counter,filename))

		if IS_GROUP_CHAT == False:
			try:
				vk.messages.send(user_id = ID ,message=MESSAGE,
							attachment = 'photo{}_{}'.format(upload_img['owner_id'],upload_img['id']),
							random_id = 0)
			except Exception as e:
				print(e)
				break
		elif IS_GROUP_CHAT == True:
			try:
				vk.messages.send(chat_id = ID, message=MESSAGE,
								attachment='photo{}_{}'.format(upload_img['owner_id'],upload_img['id']),
								random_id = 0)
			except Exception as e:
				print(e)
				break

		counter += 1

		# giving a bit of a rest to VK
		sleep(random.uniform(3.0,4.0))
		pass


if __name__ == '__main__':
	images = ''
	token = ''
	id = 0

	sendImages(img_dir = images, ID = id, IS_GROUP_CHAT = False, TOKEN = token)

	pass
