package com.xiaoyou.library.common.util

import android.graphics.Camera
import android.graphics.Matrix
import android.view.animation.Animation
import android.view.animation.Transformation

/**
 * @description 3D旋转动画效果
 * @author 小游
 * @data 2021/03/07
 */
class Rotate3dAnimation(fromDegrees: Float, toDegrees: Float,
                        centerX: Float, centerY: Float, depthZ: Float, reverse: Boolean) : Animation() {
    private val mFromDegrees: Float = fromDegrees
    private val mToDegrees: Float = toDegrees
    private val mCenterX: Float = centerX
    private val mCenterY: Float = centerY
    private val mDepthZ: Float = depthZ
    private val mReverse: Boolean = reverse
    private var mCamera: Camera? = null
    override fun initialize(width: Int, height: Int, parentWidth: Int, parentHeight: Int) {
        super.initialize(width, height, parentWidth, parentHeight)
        mCamera = Camera()
    }

    override fun applyTransformation(interpolatedTime: Float, t: Transformation) {
        val fromDegrees = mFromDegrees
        val degrees = fromDegrees + (mToDegrees - fromDegrees) * interpolatedTime
        val centerX = mCenterX
        val centerY = mCenterY
        val camera: Camera? = mCamera
        val matrix: Matrix = t.getMatrix()
        camera?.save()
        if (mReverse) {
            camera?.translate(0.0f, 0.0f, mDepthZ * interpolatedTime)
        } else {
            camera?.translate(0.0f, 0.0f, mDepthZ * (1.0f - interpolatedTime))
        }
        camera?.rotateY(degrees)
        camera?.getMatrix(matrix)
        camera?.restore()
        matrix.preTranslate(-centerX, -centerY)
        matrix.postTranslate(centerX, centerY)
    }

}